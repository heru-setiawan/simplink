package handler

import (
	"simplink/features/links/core"
	"time"
)

type DefaultResponse map[string]any

type LinkResponse struct {
	Short       string    `json:"short"`
	Destination string    `json:"desitnation"`
	ExpiredAt   time.Time `json:"expired_at"`
}

func LinkResponseFromEntity(ent core.Link) LinkResponse {
	return LinkResponse{
		Short:       ent.Short,
		Destination: ent.Destination,
		ExpiredAt:   ent.ExpiredAt,
	}
}
