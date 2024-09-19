package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm/clause"
	"project/service1/model"
	"project/service1/pg"
)

// ObjectRepository is a object repository.
type ObjectRepository struct {
	db *pg.DB
}

func (o ObjectRepository) Upsert(ctx context.Context, lastSeen *model.DBObject) error {
	if err := o.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"last_seen"}),
	}).Create(lastSeen).Error; err != nil {
		return fmt.Errorf("failed to upsert record: %w", err)
	}
	return nil
}

func (o ObjectRepository) Create(ctx context.Context, lastSeen *model.DBObject) error {

	if err := o.db.WithContext(ctx).Create(lastSeen).Error; err != nil {
		return fmt.Errorf("failed to create record: %w", err)
	}
	return nil
}

func (o ObjectRepository) getByID(ctx context.Context, id uint) (*model.DBObject, error) {
	//TODO implement me
	panic("implement me")
}

func (o ObjectRepository) Update(ctx context.Context, lastSeen *model.DBObject) error {

	if err := o.db.WithContext(ctx).Save(lastSeen).Error; err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}
	return nil
}

func (o ObjectRepository) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	if err := o.db.WithContext(ctx).Delete(&model.DBObject{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete record: %w", err)
	}
	return nil
}

// New creates a new repository.
func New(db *pg.DB) ObjectRepoInterface {
	return &ObjectRepository{
		db: db,
	}
}
