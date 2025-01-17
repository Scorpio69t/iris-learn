package environment

import (
	"os"
	"strings"
)

const (
	PROD Env = "prod"
	DEV  Env = "dev"
)

type Env string

// String returns the string representation of the environment
func (e Env) String() string {
	return string(e)
}

// ReadEnv reads the environment variable key and returns the value if it exists, otherwise it returns def
func ReadEnv(key string, def Env) Env {
	v := GetEnv(key, def.String())
	if v == "" {
		return def
	}

	env := Env(strings.ToLower(v))
	switch env {
	case PROD, DEV:
	default:
		panic("unexpected environment value: " + v)
	}

	return env
}

// GetEnv returns the value of the environment variable key if it exists, otherwise it returns def
func GetEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}
