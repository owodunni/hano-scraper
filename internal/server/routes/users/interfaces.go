package users

import (
	"context"

	"github.com/owodunni/hano-scraper/internal/models"
)

type Logger interface {
	Debugf(format string, args ...any)
	Error(s string)
}

type Processor interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByID(ctx context.Context, id uint64) (user models.User, err error)
}
