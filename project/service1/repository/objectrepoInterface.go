package repository

import (
	"context"
	"project/service1/model"
)

type ObjectRepoInterface interface {
	Create(ctx context.Context, lastSeen *model.DBObject) error
	getByID(ctx context.Context, id uint) (*model.DBObject, error)
	Update(ctx context.Context, lastSeen *model.DBObject) error
	Delete(ctx context.Context, id uint) error
	Upsert(ctx context.Context, lastSeen *model.DBObject) error
}
