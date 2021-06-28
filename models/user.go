package models

import (
	"time"
)

type User struct {
	ID        *string   `param:"id" json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Surname   *string   `json:"surname,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updateAt,omitempty"`
}
