package handler

import "simplink/features/links/core"

type LinkCreateRequest struct {
	Short       string `json:"short,omitempty"`
	Destination string `json:"destination"`
}

func (req *LinkCreateRequest) ToEntity() core.Link {
	return core.Link{
		Short:       req.Short,
		Destination: req.Destination,
	}
}
