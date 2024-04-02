package edc

import (
	"gorm.io/gorm"
)

type BaseRepository struct {
	Model interface{}
}

func (r *BaseRepository) SetModel(model interface{}) {
	r.Model = model
}

func (r *BaseRepository) DB() *gorm.DB {
	return Edc.DB.Model(r.Model)
}

func (r *BaseRepository) Count() (int64, error) {
	var count int64
	res := r.DB().Count(&count)
	return count, res.Error
}

func (r *BaseRepository) GetAll(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return r.DB().Find(dest, conds...)
}

func (r *BaseRepository) GetByID(dest interface{}, id string, deleted bool) (tx *gorm.DB) {
	q := r.DB()
	if deleted {
		q = q.Unscoped()
	}
	return q.Where("id = ?", id).First(dest)
}

func (r *BaseRepository) Paginate(
	dest interface{},
	args PaginateArgs,
	conds ...interface{},
) (tx *gorm.DB) {
	return r.DB().
		Limit(args.Limit).
		Offset((args.Page-1)*args.Limit).
		Order(args.OrderBy+" "+args.Order).
		Find(dest, conds...)
}

func (r *BaseRepository) Create(dest interface{}) (tx *gorm.DB) {
	return r.DB().Create(dest)
}

func (r *BaseRepository) Update(dest interface{}) (tx *gorm.DB) {
	return Edc.DB.Save(dest)
}

func (r *BaseRepository) Delete(dest interface{}) (tx *gorm.DB) {
	return r.DB().Delete(dest)
}

func (r *BaseRepository) Restore(dest interface{}) (tx *gorm.DB) {
	return r.DB().Unscoped().Model(dest).Update("deleted_at", nil)
}

func (r *BaseRepository) HardDelete(dest interface{}) (tx *gorm.DB) {
	return r.DB().Unscoped().Delete(dest)
}
