package model

import (
	golog "log"

	"github.com/go-sohunjug/logger"
)

var DefaultLogger logger.Logger = logger.NewStdLogger(golog.Writer())
