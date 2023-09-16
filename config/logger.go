package config

import "log"

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}
