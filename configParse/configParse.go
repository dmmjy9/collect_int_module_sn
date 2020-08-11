package configParse

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	SSHConfig	*SSHConfig	`json:"ssh"`
	DBConfig	*DBConfig	`json:"mysql"`
}

type SSHConfig struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type DBConfig struct {
	Host		string	`json:"host"`
	Port		int		`json:"port"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Database	string	`json:"database"`
}

func loadConfig(filename string) (Config, error) {
	var jsonParm Config
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	jsonValue, _ := ioutil.ReadAll(file)
	if err := json.Unmarshal([]byte(jsonValue), &jsonParm); err != nil {
		return Config{}, err
	} else {
		return jsonParm, nil
	}
}

func GetDBConfig(filename string) (DBConfig, error) {
	if jsonParm, err := loadConfig(filename); err != nil {
		return DBConfig{}, err
	} else {
		return *jsonParm.DBConfig, nil
	}
}

func GetSSHConfig(filename string) (SSHConfig, error) {
	if jsonParm, err := loadConfig(filename); err != nil {
		return SSHConfig{}, err
	} else {
		return *jsonParm.SSHConfig, nil
	}
}
