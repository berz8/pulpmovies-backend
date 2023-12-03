package validators

type Movie struct {
  ID int32 `json:"id" validate:"required"`
  ImdbID string `json:"imdb_id" validate:"required"`
  OriginalTitle string `json:"original_title" validate:"required"`
  Title string `json:"title" validate:"required"`
  OriginalLanguage string `json:"original_language" validate:"required"`
  Overview string `json:"overview" validate:"required"`
  PosterPath string `json:"poster_path" validate:"required"`
  BackdropPath string `json:"backdrop_path" validate:"required"`
  ReleaseDate string `json:"release_date" validate:"required"`
}
