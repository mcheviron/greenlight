package data

import (
	"time"

	"github.com/mcheviron/greenlight/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id,omitempty"`
	CreatedAt time.Time `json:"-"` // When the movie was added to the database
	Title     string    `json:"title,omitempty"`
	Year      int32     `json:"year,omitempty"`    // The year the movie was released
	Runtime   Runtime   `json:"runtime,omitempty"` // Movie runtime in minutes
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version,omitempty"` // The version number starts at 1 and will be incremented each time the movie information is updated
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greated than 1888")
	v.Check(
		movie.Year <= int32(time.Now().Year()),
		"year",
		"must not be in the future",
	)

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	// NOTE: avoid using len() because it can potentially cause a panic
	// v.Check(len(movie.Genres) > 0, "genres", "must be provided")
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
