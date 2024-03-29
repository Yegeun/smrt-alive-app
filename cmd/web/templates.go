package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/Yegeun/smrt-alive-app/pkg/forms"
	"github.com/Yegeun/smrt-alive-app/pkg/models"
)

func humanDate(t time.Time) string {
	return  t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error){
	//Initialize the new map to act as the cashe.
	cache := map[string]*template.Template{}

	//Use the filepat.glob function to get a slice of all file paths with
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))

	if err != nil {
		return nil, err
	}
	// Loop throught all the pages one by one

	for _, page := range pages{
		// Extrace the file name and assign it to new variables

		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before you
		// call the ParseFiles() method. This means we have to use template.New() to
		// create an empty template set, use the Funcs() method to register the
		// template.FuncMap, and then parse the file as normal.
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
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
	CurrentYear int
	Form *forms.Form
	Snippet *models.Snippet
	Snippets []*models.Snippet
}