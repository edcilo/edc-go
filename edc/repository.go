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

func (r *BaseRepository) Count(args RepositoryCountArgs) *gorm.DB {
	return r.DB().Count(&args.Total)
}

func (r *BaseRepository) GetAll(args RepositoryGetAllArgs, conds ...interface{}) (tx *gorm.DB) {
	return r.DB().Order(args.Order).Find(args.Dest, conds...)
}

func (r *BaseRepository) GetByID(args RepositoryGetByIDArgs) (tx *gorm.DB) {
	q := r.DB()
	if args.Deleted {
		q = q.Unscoped()
	}
	return q.Where("id = ?", args.ID).First(args.Dest)
}

func (r *BaseRepository) Paginate(
	args RepositoryPaginateArgs,
	conds ...interface{},
) (tx *gorm.DB) {
	return r.DB().
		Limit(args.Limit).
		Offset((args.Page-1)*args.Limit).
		Order(args.Order).
		Find(args.Dest, conds...)
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
