package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"log"
	"os"
	"reflect"
	"sync"
)

var db *sqlx.DB
var once sync.Once

// GetEnv : is a function to get system environment
func GetEnv(key string) string {
	r := os.Getenv(key)

	if r == "" {
		if _, ok := defaultConfig[key]; !ok {
			return ""
		}
		r = defaultConfig[key]
	}

	return r
}

// GetDB : is a function to get Database Connection
func GetDB() *sqlx.DB {
	once.Do(func() {
		var err error

		conn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			GetEnv(DB_HOST),
			GetEnv(DB_PORT),
			GetEnv(DB_USER),
			GetEnv(DB_PASS),
			GetEnv(DB_NAME),
		)

		db, err = sqlx.Connect("postgres", conn)

		if nil != err {
			log.Fatal(err)
		}
	})
	return db
}

// GetQueryConfig is a function to get default fazzdb config
func GetQueryConfig() fazzdb.Config {
	return GetIfQueryConfig(I_QUERY_CONFIG)
}

// GetIfQueryConfig get config as fazzdb.Config
func GetIfQueryConfig(key string) fazzdb.Config {
	var t fazzdb.Config
	return getIf(key, t).(fazzdb.Config)
}

func getIf(key string, p interface{}) interface{} {
	t := reflect.TypeOf(p)

	if _, ok := baseInterface[key]; !ok {
		panic(fmt.Sprintf(`config with key "%s" not found`, key))
	}

	if keyType := reflect.TypeOf(baseInterface[key]); t != keyType {
		panic(fmt.Sprintf(`different type of config with key "%s" got %s expected %s`, key, keyType, t))
	}

	return baseInterface[key]
}
