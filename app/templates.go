package app

import (
	"html/template"

	bin "github.com/gophergala/golang-sizeof.tips/internal/bindata/templates"
)

const templatesDir = "templs/"

var templates map[string]*template.Template

func prepareTemplates() error {
	templates = make(map[string]*template.Template)
	baseData, err := bin.Asset(templatesDir + "parts/base.tmpl")
	if err != nil {
		return err
	}
	for _, name := range []string{
		"index",
	} {
		assetData, err := bin.Asset(templatesDir + name + ".tmpl")
		if err != nil {
			return err
		}
		templates[name], err = template.New(name).Parse(
			string(baseData) + string(assetData),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
