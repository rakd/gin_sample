package config

import (
	"fmt"
	"log"
	"os"
)

var envChecked bool

const (
	envJWTSaltName = "JWT_SALT"

	// it's better NOT to change it. if change it, all user's sessoins are expired.
	additionalJWTSalt = "etadpWjainL3xsqU7FVZbVYkQmeua"
	//
)

var jwtSalt string

func init() {
	CheckEnvs()
}

// CheckEnvs ...
func CheckEnvs() error {
	if envChecked {
		return nil
	}

	envJWTSalt := os.Getenv(envJWTSaltName)

	if envJWTSalt == "" {
		return fmt.Errorf("you need to set env. (%s)", envJWTSaltName)
	}

	jwtSalt = envJWTSalt + additionalJWTSalt
	envChecked = true
	return nil
}

// GetJWTSalt ...
func GetJWTSalt() string {
	if envChecked == false {
		log.Print("NO CONFIG ENVS CHECKED")
		os.Exit(1)
	}
	return jwtSalt
}
