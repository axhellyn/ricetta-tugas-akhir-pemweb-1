package detailcontroller

import (
	"html/template"
	"net/http"
	"ricetta/entities"
	"ricetta/models/detailmodel"
	"ricetta/models/resepmodel"
	"strconv"
)

func DetailRecipe(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if idStr == "" {
        http.Error(w, "Recipe ID not provided", http.StatusBadRequest)
        return
    }

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	resep, err := resepmodel.GetResepByID(id)
    if err != nil {
        http.Error(w, "Recipe not found", http.StatusNotFound)
        return
    }

	bahans, err := detailmodel.GetBahanByResepID(id)
    if err != nil {
        http.Error(w, "Recipe not found", http.StatusNotFound)
        return
    }

	data := struct {
        Resep  entities.Resep
        Bahans []entities.Bahan
    }{
        Resep:  resep,
        Bahans: bahans,
    }

	temp := template.Must(template.ParseFiles("views/detail/detail-recipe.html"))
	temp.Execute(w, data)
}