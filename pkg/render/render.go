package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/dh-n/go-course/pkg/config"
	"github.com/dh-n/go-course/pkg/models"
)

var functions template.FuncMap
var app *config.AppConfig

func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData){
	var tc map[string]*template.Template
	if app.UseCache{
		tc = app.TemplateCache
	}else {
		tc,_ = CacheTemplate()
	}
	// parsedTemplates, err := template.ParseFiles("templates/"+tmpl)
	parsedTemplate, ok := tc[tmpl]
	if !ok {
		log.Fatal("error parsing template")
	}

	 buf := new(bytes.Buffer)


	 td = AddDefaultData(td)
	// err := parsedTemplate.Execute(w, nil)
	err := parsedTemplate.Execute(buf, td)

	if err != nil{
		fmt.Println("Error parsing template ", err)
	}

	_,err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser")
	}
}

func CacheTemplate()(map[string]*template.Template, error){
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages{
		filename := filepath.Base(page)
		fmt.Println("Currently building ", filename)

		//build the template set
		ts,err := template.New(filename).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches,err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts,err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[filename] = ts
	}
	return myCache, nil
}