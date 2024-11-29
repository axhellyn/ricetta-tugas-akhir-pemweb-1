package resepmodel

import (
	"ricetta/config"
	"ricetta/entities"
)

func GetAll() []entities.Resep {
	rows, err := config.DB.Query("SELECT * FROM resep")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var reseps []entities.Resep

	for rows.Next(){
		var resep entities.Resep
		if err := rows.Scan(&resep.Id_resep, &resep.Nama, &resep.Img_url, &resep.Details, &resep.Created_at); err != nil{
			panic(err)
		}

		reseps = append(reseps, resep)
	}

	return reseps
}