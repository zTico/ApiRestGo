package models

type Personalitie struct {
	Id      int    `json:"id"`
	Name    string `json:"nome"`
	History string `json:"historia"`
}

var Personalities []Personalitie
