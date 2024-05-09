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
	return r.DB().Count(args.Total)
}

func (r *BaseRepository) GetAll(args RepositoryGetAllArgs, conds ...interface{}) (tx *gorm.DB) {
	q := r.DB()
	for _, preload := range args.Preload {
		q = q.Preload(preload)
	}
	return q.Order(args.Order).Find(args.Dest, conds...)
}

func (r *BaseRepository) GetByID(args RepositoryGetByIDArgs) (tx *gorm.DB) {
	q := r.DB()
	for _, preload := range args.Preload {
		q = q.Preload(preload)
	}
	if args.Deleted {
		q = q.Unscoped()
	}
	return q.Where("id = ?", args.ID).First(args.Dest)
}

func (r *BaseRepository) GetByIDs(
	args RepositoryGetByIDsArgs,
	conds ...interface{},
) (tx *gorm.DB) {
	q := r.DB()
	for _, preload := range args.Preload {
		q = q.Preload(preload)
	}
	if args.Deleted {
		q = q.Unscoped()
	}
	return q.
		Where("id IN (?)", args.IDs).
		Find(args.Dest, conds...)
}

func (r *BaseRepository) Paginate(
	args RepositoryPaginateArgs,
	conds ...interface{},
) (tx *gorm.DB) {
	q := r.DB()
	for _, preload := range args.Preload {
		q = q.Preload(preload)
	}
	return q.
		Limit(args.Limit).
		Offset((args.Page-1)*args.Limit).
		Order(args.Order).
		Find(args.Dest, conds...)
}

func (r *BaseRepository) PaginateMetadata(args RepositoryPaginateMetadataArgs) error {
	var total int64
	countArgs := RepositoryCountArgs{Total: &total}
	if res := r.Count(countArgs); res.Error != nil {
		return res.Error
	}

	lastPage := int(total)/args.Limit + 1
	beforePage := args.Page - 1
	nextPage := args.Page + 1

	(*args.Metadata)["current"] = args.Page
	(*args.Metadata)["last"] = lastPage
	(*args.Metadata)["total"] = total

	if beforePage > 0 {
		(*args.Metadata)["previous"] = beforePage
	} else {
		(*args.Metadata)["previous"] = nil
	}

	if nextPage <= lastPage {
		(*args.Metadata)["next"] = nextPage
	} else {
		(*args.Metadata)["next"] = nil
	}

	return nil
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

func (r *BaseRepository) ClearRelations(dest interface{}, relation string) error {
	return r.DB().Model(dest).Association(relation).Clear()
}
