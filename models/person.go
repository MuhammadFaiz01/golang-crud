package models

type Person struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
	BirthDate string `json:"birth_date"`
	Address  string `json:"address"`
}