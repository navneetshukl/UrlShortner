package render

import (
	"html/template"
	"net/http"
	"urlshortner/models"
)

func RenderTemplate(w http.ResponseWriter, url string, td *models.TemplateData) {
	tmpl, _ := template.ParseFiles("./templates/" + url)

	tmpl.Execute(w, td)
}
