package edc

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
