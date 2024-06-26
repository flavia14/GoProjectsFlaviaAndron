package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("../templates/" + tmpl) 
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

var tc = make(map[string] *template.Template)

func RenderTemplateTest(w http.ResponseWriter. t string) {
	var tmpl *template.Template
	var err error

	//check to see if we already have the template in our chace
	_, inMap := tc[t]
	if !inMap {
		//need to create the template
		err = createTemplateChache(t)
		if err != nil {
			log.Println(err)
		}

	} else {
		log,Println("using chached tamplate")
	}
	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
}

func createTemplateChache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), 
		"./templates/base.layout.tmpl",
	}
	
	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	
	if err != nil {
		return err
	}

	// add template to cache (map)

	tc[t] = tmpl

	return nil
}

