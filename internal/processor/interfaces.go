package processor

import (
	"context"

	"github.com/owodunni/hano-scraper/internal/models"
)

type Database interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByID(ctx context.Context, id uint64) (user models.User, err error)
}
