package edc

import "gorm.io/gorm"

type CacheEngine string

const (
	Redis CacheEngine = "redis"
)

type CacheDSN struct {
	Host     string
	Port     int
	User     string
	Password string
	Database int
}

type CacheSetupArgs struct {
	Engine CacheEngine
	DSN    CacheDSN
}

type DBEngine string

const (
	SQLite   DBEngine = "sqlite"
	Postgres DBEngine = "postgres"
	MySQL    DBEngine = "mysql"
)

type DBDSN struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type DBSetupArgs struct {
	Engine DBEngine
	DSN    DBDSN
}

type NewEDCArgs struct {
	DB    DBSetupArgs
	Cache CacheSetupArgs
}

type PaginateArgs struct {
	Page    int
	Limit   int
	OrderBy string
	Order   string
}

type BaseRepositoryInterface interface {
	SetModel(model interface{})
	DB() *gorm.DB
	Count() (int64, error)
	GetAll(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	GetByID(dest interface{}, id string, deleted bool) (tx *gorm.DB)
	Paginate(dest interface{}, args PaginateArgs, conds ...interface{}) (tx *gorm.DB)
	Create(dest interface{}) (tx *gorm.DB)
	Update(dest interface{}) (tx *gorm.DB)
	Delete(dest interface{}) (tx *gorm.DB)
}

// Configuration struct
type ConfigurationApp struct {
	Name        string
	Description string
	Version     string
	Host        string
	Port        int
}

type ConfigurationDB struct {
	Engine   DBEngine
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type ConfigurationCache struct {
	Engine   CacheEngine
	Host     string
	Port     int
	User     string
	Password string
	Database int
}

type Configuration struct {
	App   ConfigurationApp
	DB    ConfigurationDB
	Cache ConfigurationCache
}
