package edc

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
	DB DBSetupArgs
}
