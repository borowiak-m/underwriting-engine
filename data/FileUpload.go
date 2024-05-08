package data

import (
	"time"

	"github.com/google/uuid"
)

type FileType int

const (
	FileTypeSubscription FileType = iota
)

type FileUpload struct {
	ID         uuid.UUID      `json:"id"`
	CustomerID uuid.UUID      `json:"customerId"`
	Type       int            `json:"type"`
	Mapping    map[string]int `json:"mapping"`
	CreatedAt  time.Time      `json:"createdAt"`
}
