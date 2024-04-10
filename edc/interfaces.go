package edc

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

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

type DBSeederFunc func(*gorm.DB) error

type NewEDCArgs struct {
	DB    DBSetupArgs
	Cache CacheSetupArgs
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

// Repository structs
type BaseRepositoryInterface interface {
	SetModel(model interface{})
	DB() *gorm.DB
	Count(args RepositoryCountArgs) (tx *gorm.DB)
	GetAll(args RepositoryGetAllArgs, conds ...interface{}) (tx *gorm.DB)
	GetByID(args RepositoryGetByIDArgs) (tx *gorm.DB)
	Paginate(args RepositoryPaginateArgs, conds ...interface{}) (tx *gorm.DB)
	Create(dest interface{}) (tx *gorm.DB)
	Update(dest interface{}) (tx *gorm.DB)
	Delete(dest interface{}) (tx *gorm.DB)
}

type RepositoryGetAllArgs struct {
	Dest    interface{}
	Order   string
	Preload []string
}

type RepositoryGetByIDArgs struct {
	Dest    interface{}
	ID      string
	Deleted bool
	Preload []string
}

type RepositoryGetByIDsArgs struct {
	Dest    interface{}
	IDs     []string
	Deleted bool
	Preload []string
}

type RepositoryPaginateArgs struct {
	Dest    interface{}
	Page    int
	Limit   int
	Order   string
	Preload []string
}

type RepositoryCountArgs struct {
	Total *int64
}

type RepositoryPaginateMetadataArgs struct {
	Page     int
	Limit    int
	Metadata *fiber.Map
}
