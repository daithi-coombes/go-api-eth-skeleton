package utils

import (
	"os"
	"path"
	"text/template"
)

// GetTemplate will load a html/template instance.
func GetTemplate(templateName string) (*template.Template, error) {
	var tmpl *template.Template

	dir, err := os.Getwd()
	if err != nil {
		return tmpl, err
	}

	filepath := path.Join(dir, "../../web/"+templateName+".html")
	tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		return tmpl, err
	}

	return tmpl, nil
}
