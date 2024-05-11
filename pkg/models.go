package helper

import "github.com/google/uuid"

type Player struct {
	ID        uuid.UUID `json:"id"`
	ClubId    uuid.UUID `json:"clubId"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	ShirtName string    `json:"shirtName"`
}
