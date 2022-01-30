package driver

import (
	"os"
)

type FileDriver struct {
	file      os.File
	connected bool
}

func (driver *FileDriver) Connect() error {
	if driver.connected {
		return nil
	}

	driver.connected = true

	return nil
}

func (driver *FileDriver) Disconnect() {
	driver.connected = false
}

func (driver *FileDriver) IsConnected() bool {
	return driver.connected
}
