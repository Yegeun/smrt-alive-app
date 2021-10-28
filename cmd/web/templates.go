package main

import (
	"github.com/Yegeun/smrt-alive-app/pkg/models"
	"html/template"
	"path/filepath"
)

func newTemplateCache(dir string) (map[string]*template.Template, error){
	//Initialize the new map to act as the cashe.
	cache := map[string]*template.Template{}

	//Use the filepat.glob function to get a slice of all file paths with
	// extension

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))

	if err != nil {
		return nil, err
	}

	// Loop throught all the pages one by one

	for _, page := range pages{
		// Extrace the file name and assign it to new variables

		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// use the parseglob method to add any layout templates
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		//use the parseglob method add any partial templates
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	//return to the map
	return  cache, nil
}

//Define a templateData type to act as the holding structure for
//any dynamic data that we want to pass to our HTML templates.
//At the moment it only contains one field, but we'll add more
//to it as the build progresses.
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}