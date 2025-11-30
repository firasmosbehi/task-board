// Package main defines the data models used by the TaskBoard backend service.
package main

import "time"

// Task represents a task item stored in the database.
// It includes metadata fields automatically managed by GORM.
type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
