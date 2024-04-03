package edc

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct{}

func (e *Env) Load(filenames ...string) {
	godotenv.Load(filenames...)
}

func (e *Env) Get(key string, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}

func (e *Env) GetBool(key string, defaultVal bool) bool {
	if val, exists := os.LookupEnv(key); exists {
		return val == "true" || val == "1"
	}
	return defaultVal
}

func (e *Env) GetInt(key string, defaultVal int) int {
	if val, exists := os.LookupEnv(key); exists {
		if valInt, err := strconv.Atoi(val); err == nil {
			return valInt
		}
	}
	return defaultVal
}
