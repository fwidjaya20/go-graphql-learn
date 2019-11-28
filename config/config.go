package config

import (
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"time"
)

// noinspection ALL
const (
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
	DB_NAME = "DB_NAME"
	DB_USER = "DB_USER"
	DB_PASS = "DB_PASS"
)

// noinspection ALL
const (
	HTTP_ADDR = "HTTP_ADDRESS"
)

// noinspection ALL
const (
	I_QUERY_CONFIG         = "QUERY_CONFIG"
	I_READ_TIMEOUT         = "READ_TIMEOUT"
	I_WRITE_TIMEOUT        = "WRITE_TIMEOUT"
	I_IDLE_TIMEOUT         = "IDLE_TIMEOUT"
	I_WAIT_SHUTDOWN        = "WAIT_SHUTDOWN"
	I_CREDENTIAL_EXPIRE    = "CREDENTIAL_EXPIRE"
	I_REDIS_SESSION_EXPIRE = "REDIS_SESSION_EXPIRE"
	I_THROTTLE_LIMIT       = "THROTTLE_LIMIT"
	I_THROTTLE_DURATION    = "THROTTLE_DURATION"
	I_THROTTLE_TYPE        = "THROTTLE_TYPE"
)

var defaultConfig = map[string]string{
	// Database Configuration
	DB_HOST: "localhost",
	DB_PORT: "5432",
	DB_NAME: "fazzlearning_auth",
	DB_USER: "postgres",
	DB_PASS: "postgres",

	// Transport Configuration
	HTTP_ADDR: ":8000",
}

var baseInterface = map[string]interface{}{
	// I_QUERY_CONFIG conf for DB limit, offset and lock
	I_QUERY_CONFIG: fazzdb.Config{
		Limit:           0,
		Offset:          0,
		Lock:            fazzdb.LO_NONE,
		DevelopmentMode: true,
	},
	// I_READ_TIMEOUT read timeout
	I_READ_TIMEOUT: 5 * time.Minute,
	// I_WRITE_TIMEOUT write timeout
	I_WRITE_TIMEOUT: 5 * time.Minute,
	// I_IDLE_TIMEOUT idle timeout
	I_IDLE_TIMEOUT: 30 * time.Second,
	// I_WAIT_SHUTDOWN graceful shutdown delay
	I_WAIT_SHUTDOWN: 5 * time.Second,
	// I_CREDENTIAL_EXPIRE credential expired time
	I_CREDENTIAL_EXPIRE: 24 * 25 * time.Hour,
	// I_REDIS_SESSION_EXPIRE const for redis expired session
	I_REDIS_SESSION_EXPIRE: 24 * time.Hour,
	// I_THROTTLE_LIMIT limit per hit for throttling
	I_THROTTLE_LIMIT: 40,
	// I_THROTTLE_DURATION duration for throttle TTL
	I_THROTTLE_DURATION: time.Minute,
	// I_THROTTLE_TYPE type checker for throttle
	I_THROTTLE_TYPE: 0,
}