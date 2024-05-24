package models

import "gorm.io/gorm"

// Структура задачи
type Task struct {
	gorm.Model
	ID           string   `json:"id" gorm:"primaryKey"`
	Description  string   `json:"description" gorm:"text"`
	Note         string   `json:"note" gorm:"text"`
	Applications []string `json:"applications"`
}
