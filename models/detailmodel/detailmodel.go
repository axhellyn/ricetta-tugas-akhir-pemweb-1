package detailmodel

import (
	"ricetta/config"
	"ricetta/entities"
)

func GetBahanByResepID(resepID int) ([]entities.Bahan, error) {
	var bahans []entities.Bahan
	query := "SELECT * FROM bahan WHERE id_resep = ?"
	rows, err := config.DB.Query(query, resepID)

	if err != nil {
        return nil, err
    }

	defer rows.Close()

	for rows.Next() {
		var bahan entities.Bahan
		if err := rows.Scan(&bahan.Id, &bahan.Id_resep, &bahan.Name); err != nil {
			panic(err)
		}

		bahans = append(bahans, bahan)
	}

	return bahans, nil
}