package models

type Personalitie struct {
	Id      int    `json:"id" gorm:"column:id"`
	Name    string `json:"nome" gorm:"column:nome"`
	History string `json:"historia" gorm:"column:historia"`
}

func (p *Personalitie) TableName() string {
	return "personalidades"
}

var Personalities []Personalitie
