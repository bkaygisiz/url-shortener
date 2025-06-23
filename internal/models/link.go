package models

import "time"

// Link représente un lien raccourci dans la base de données.
// Les tags `gorm:"..."` définissent comment GORM doit mapper cette structure à une table SQL.
// ID qui est une primaryKey
// Shortcode : doit être unique, indexé pour des recherches rapide (voir doc), taille max 10 caractères
// LongURL : doit pas être null
// CreateAt : Horodatage de la créatino du lien

type Link struct {
	ID        uint      `gorm:"primaryKey"`                   // ID qui est une primaryKey
	ShortCode string    `gorm:"uniqueIndex;size:10;not null"` // Shortcode : doit être unique, indexé pour des recherches rapide, taille max 10 caractères
	LongURL   string    `gorm:"not null"`                     // LongURL : doit pas être null
	CreatedAt time.Time // Horodatage de la création du lien
	UpdatedAt time.Time // Horodatage de la dernière mise à jour (GORM l'ajoute automatiquement)
}
