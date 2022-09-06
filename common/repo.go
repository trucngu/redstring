package common

import (
	"gorm.io/gorm"
)

type Repository[T, ID any] struct {
	Db *gorm.DB
}

func (r *Repository[T, ID]) Create(entity *T) error {
	res := r.Db.Create(entity)
	return res.Error
}

func (r *Repository[T, ID]) Get() ([]*T, error) {
	var entities []*T
	res := r.Db.Find(&entities)
	return entities, res.Error
}

func (r *Repository[T, ID]) GetById(id ID) (*T, error) {
	var entity *T
	res := r.Db.First(&entity, id)
	return entity, res.Error
}

func (r *Repository[T, ID]) Delete(id ID) error {
	var entity *T
	return r.Db.Delete(&entity, id).Error
}

func (r *Repository[T, ID]) Update(id ID, mapper func(*T)) (*T, error) {
	var entity *T
	res := r.Db.First(&entity, id)
	mapper(entity)
	r.Db.Save(&entity)
	return entity, res.Error
}
