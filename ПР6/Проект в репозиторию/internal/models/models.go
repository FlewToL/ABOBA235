package models

import "time"

// User - автор заметок: один пользователь может иметь много заметок.
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:200;uniqueIndex;not null"`
	Notes     []Note
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Note - заметка, принадлежащая пользователю и связанная с тегами.
type Note struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:200;not null"`
	Content   string `gorm:"type:text"`
	UserID    uint   `gorm:"not null"`
	User      User
	Tags      []Tag `gorm:"many2many:note_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Tag - тег, который может использоваться в разных заметках.
type Tag struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:50;uniqueIndex;not null"`
	Notes     []Note `gorm:"many2many:note_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}