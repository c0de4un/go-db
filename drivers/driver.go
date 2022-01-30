package driver

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Configuration struct {
	uri      string
	name     string
	login    string
	password string
}

type Row struct {
	cols map[string]string
}

type ISerializable interface {
	Serialize() string
}

type IDeserializable interface {
	Deserialize()
}

type IDriver interface {
	Write(serializable *ISerializable) error
	Read(deserializable *IDeserializable) error
}

func (cfg *Configuration) SetUri(uri string) {
	cfg.uri = uri
}

func (cfg *Configuration) SetName(name string) error {
	if len(name) < 2 {
		return errors.New("Configuration::SetName: Invalid name")
	}
	cfg.name = name

	return nil
}

func (cfg *Configuration) GetName() string {
	return cfg.name
}

func (cfg *Configuration) SetLogin(login string) {
	cfg.login = login
}

func (cfg *Configuration) SetPassword(password string) {
	cfg.password = password
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

func LoadConfig(path string) (Configuration, error) {
	var output Configuration

	file, err := os.Open(path)
	if err != nil {
		return output, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string

	// Uri
	scanner.Scan()
	input = strings.TrimSpace(scanner.Text())
	input = strings.Replace(input, "uri=", "", 1)
	input = strings.Replace(input, "\"", "", 2)
	output.SetUri(input)

	// Name
	scanner.Scan()
	input = strings.ReplaceAll(scanner.Text(), " ", "")
	input = strings.Replace(input, "name=", "", 1)
	input = strings.Replace(input, "\"", "", 2)
	output.SetName(input)

	// Login
	scanner.Scan()
	input = strings.ReplaceAll(scanner.Text(), " ", "")
	input = strings.Replace(input, "login=", "", 1)
	input = strings.Replace(input, "\"", "", 2)
	output.SetLogin(input)

	// Password
	scanner.Scan()
	input = strings.ReplaceAll(scanner.Text(), " ", "")
	input = strings.Replace(input, "password=", "", 1)
	input = strings.Replace(input, "\"", "", 2)
	output.SetPassword(input)

	return output, nil
}
