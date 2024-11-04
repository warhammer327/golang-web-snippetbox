package main

import (
	"html/template"
	"path/filepath"
	"warhammer327.github.io/snippetbox/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
