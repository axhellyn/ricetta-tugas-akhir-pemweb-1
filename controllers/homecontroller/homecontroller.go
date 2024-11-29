package homecontroller

import (
	"html/template"
	"net/http"
	"ricetta/models/resepmodel"
)

func Welcome(w http.ResponseWriter, r *http.Request ){
	temp, err := template.ParseFiles("views/home/Index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func Index(w http.ResponseWriter, r *http.Request ){
	reseps := resepmodel.GetAll()
	data := map[string]any{
		"reseps": reseps,
	}

	temp, err := template.ParseFiles("views/home/index.html")
		if err != nil {
			panic(err)
		}
		
	temp.Execute(w, data)

}