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
		if err := rows.Scan(&resep.Id_resep, &resep.Nama, &resep.Img_url, &resep.Details, &resep.Created_at, &resep.Cara , &resep.Porsi); err != nil{
			panic(err)
		}

		reseps = append(reseps, resep)
	}

	return reseps
}

func GetResepByID(id int) (entities.Resep, error) {
    var resep entities.Resep
    query := "SELECT * FROM resep WHERE id_resep = ?"
    row := config.DB.QueryRow(query, id)
    err := row.Scan(&resep.Id_resep, &resep.Nama, &resep.Img_url, &resep.Details, &resep.Created_at, &resep.Cara , &resep.Porsi)
    if err != nil {
        return entities.Resep{}, err
    }
    return resep, nil
}