package repository

import (
	"context"
	"fmt"
	"project/service1/model"
	"project/service1/pg"
)

// ObjectRepository is a object repository.
type ObjectRepository struct {
	db *pg.DB
}

func (o ObjectRepository) CustomUpsert(ctx context.Context, lastSeen *model.DBObject) error {
	/*_, err := o.db.ExecContext(ctx, `INSERT INTO storage_schema.objects (id, last_seen) VALUES (?, ?)`, lastSeen.ID, lastSeen.LastSeen)
	return err*/
	panic("implement me")
}

func (o ObjectRepository) Upsert(ctx context.Context, lastSeen *model.DBObject) error {
	/*_, err := o.db.WithContext(ctx).Model(lastSeen).OnConflict("(id) DO UPDATE").Set("last_seen = EXCLUDED.last_seen").Insert()
	if err != nil {
		return fmt.Errorf("failed to upsert record: %w", err)
	}
	return err*/
	panic("implement me")
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

	/*_, err := o.db.WithContext(ctx).Model(lastSeen).Where("id = ?", lastSeen.ID).Update()
	if err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil*/
	panic("implement me")
}

func (o ObjectRepository) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

// New creates a new repository.
func New(db *pg.DB) ObjectRepoInterface {
	return &ObjectRepository{
		db: db,
	}
}
