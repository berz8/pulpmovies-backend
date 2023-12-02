package handlers

import (
	"context"
	"database/sql"
	"net/http"
  "strings"

	"github.com/berz8/pulpmovies-backend/database"
	"github.com/berz8/pulpmovies-backend/models"
	"github.com/berz8/pulpmovies-backend/validators"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/oauth2/v2"
)
type AuthGoogleResponse struct {
  Message string `json:"message"`
  Result models.User `json:"result"`
  Token string `json:"token"`
}

func AuthGoogle(c *fiber.Ctx) error {
  authBody := new(validators.AuthGoogle)

  if err := c.BodyParser(authBody); err != nil {
    return err
  }

  err, errMsgs := validators.Valid.Validate(authBody)
  if err != nil {
    return &fiber.Error{
      Code:    fiber.ErrBadRequest.Code,
      Message: strings.Join(errMsgs, " and "),
    }
  }

  tokenInfo, err := getTokenInfo(authBody.IdToken)
  if err != nil {
    return err
  }

  db:= database.DB
  user := models.User{}
  err = models.GetUserByEmail(&user, db, tokenInfo.Email)
  if err != nil {
    if err == sql.ErrNoRows {
      _, err := db.ExecContext(
        context.Background(),
        `INSERT INTO users (email, username, full_name, onboarding, account_status) VALUES (?, ?, ?, 0, 'active')`,
        tokenInfo.Email,
        tokenInfo.Email,
        authBody.FullName,
      )
      if err != nil {
        return fiber.NewError(
          fiber.StatusBadRequest,
          "Something went wrong while searching user" + err.Error(),
        )
      }
      err = models.GetUserByEmail(&user, db, tokenInfo.Email)
      // Create default Watchlsit for the new user
      _, err = db.ExecContext(
        context.Background(),
        `INSERT INTO watchlists (name, user_id, public, is_default) VALUES ('watchlist', ?, 1, 1)`,
        user.ID,
      )
      if err != nil {
        return fiber.NewError(
          fiber.StatusBadRequest,
          "Something went wrong while creating user's watchlist " + err.Error(),
        )
      }
    } else {
      return fiber.NewError(
        fiber.StatusNotFound,
        "User not found",
      )
    }
  }

  token, err := models.CreateToken(int64(user.ID))  
  if err != nil {
    return fiber.NewError(
      fiber.StatusNotFound,
      "User not found" + err.Error(),
    )
  }
  

  return c.JSON(AuthGoogleResponse{
    Result: user, 
    Token: token,
    Message: "Successfully logged in",
  }) 
}

func getTokenInfo(idToken string) (*oauth2.Tokeninfo, error) {
  oauth2Service, err := oauth2.New(&http.Client{})
  if err != nil {
    return nil, err
  }
  tokenInfoCall := oauth2Service.Tokeninfo()
  tokenInfoCall.IdToken(idToken)
  return tokenInfoCall.Do()
}
