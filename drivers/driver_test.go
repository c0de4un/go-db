package driver

import (
	"testing"
)

func TestLoad(t *testing.T) {
	config, err := LoadConfig("../driver_config_example.cfg")
	if err != nil {
		t.Errorf("database::driver::LoadConfig: %s", err)
	}

	if config.GetName() != "testDB" {
		t.Errorf("database::driver::LoadConfig: bad deserialization logic, invalid name: '%s'", config.GetName())
	}
}
