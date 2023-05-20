package repositories

import "github.com/yonisaka/idempotency/internal/entity"

type UserRepo interface {
	FindAll() ([]entity.User, error)
}
