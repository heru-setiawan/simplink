package repository

import (
	"simplink/features/links/core"
	"time"
)

func newLinkFromEntity(ent core.Link) Link {
	return Link{
		Short:       ent.Short,
		Destination: ent.Destination,
		ExpiredAt:   ent.ExpiredAt,
	}
}

type Link struct {
	Short       string `gorm:"type:varchar(16); not null; primaryKey;"`
	Destination string `gorm:"type:text;"`
	ExpiredAt   time.Time
}

func (mod *Link) ToEntity() *core.Link {
	return &core.Link{
		Short:       mod.Short,
		Destination: mod.Destination,
		ExpiredAt:   mod.ExpiredAt,
	}
}
