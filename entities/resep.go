package entities

import "time"

type Resep struct {
	Id_resep   uint
	Nama       string
	Img_url    string
	Details    string
	Created_at time.Time
	Cara string
	Porsi int
}