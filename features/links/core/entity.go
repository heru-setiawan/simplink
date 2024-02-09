package core

import (
	"math/rand"
	"time"
)

type Link struct {
	Short       string
	Destination string
	ExpiredAt   time.Time
}

func (ent *Link) GenerateShort(length int) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = charset[seededRand.Intn(len(charset))]
	}
	ent.Short = string(bytes)
}

func (ent *Link) SetExpired(duration time.Duration) {
	ent.ExpiredAt = time.Now().Add(duration)
}
