package models

type Personalitie struct {
	Name    string `json:"nome"`
	History string `json:"historia"`
}

var Personalities []Personalitie
