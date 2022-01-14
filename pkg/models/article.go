package models

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// Article id holds when an article is newly created
type ArticleId struct {
	Id int `json:"id"`
}

// Article data struct inside the DB
type Article struct {
	Title   string `json:"title" db:"title" validate:"required,min=3,max=32"`
	Content string `json:"content" db:"content" validate:"required,min=3,max=32"`
	Author  string `json:"author" db:"author" validate:"required,min=3,max=32"`
}

// Error response struct after doing struct validations
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// Remove spaces
func (a *Article) TrimSpaces() {
	a.Author = strings.TrimSpace(a.Author)
	a.Content = strings.TrimSpace(a.Content)
	a.Title = strings.TrimSpace(a.Title)
}

// Validate if fields are filled.
func (a *Article) Validate() []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(a)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse

			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errors = append(errors, &element)
		}

		return errors
	}
	return nil
}
