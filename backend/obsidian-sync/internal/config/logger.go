package config

import (
	. "github.com/savabush/lib/pkg/logging"
)

var Logger = MakeLogger(Settings.LOGGING.FILE_PATH, true)
