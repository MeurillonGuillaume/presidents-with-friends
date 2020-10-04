package player

import (
	"github.com/google/uuid"
)

type Player interface {
	Join(uuid.UUID)
}
