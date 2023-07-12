package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	LongURL  string `json:"long_url"`  // original url that user wants to short
	ShortURL string `json:"short_url"` // Shorten URL

}

type TemplateData struct{
	StringMap map[string]string
}