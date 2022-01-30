package database

import "fmt"

const DRIVER_TYPE_FILE = 1

type Row struct {
	cols map[string]string
}

type ISerializable interface {
	Serialize() string
}

type IDeserializable interface {
	Deserialize()
}

func (row *Row) GetCol(name string) (string, error) {
	output, ok := row.cols[name]
	if !ok {
		return "", fmt.Errorf("database::driver::GetCol: column '%s' is not found", name)
	}

	return output, nil
}

func (row *Row) SetCol(name string, val string) {
	row.cols[name] = val
}
