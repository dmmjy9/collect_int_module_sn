package configParse

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	SSHConfig   *SSHConfig   `json:"ssh"`
	DBConfig    *DBConfig    `json:"mysql"`
	ApiConfig   *ApiConfig   `json:"api"`
	RedisConfig *RedisConfig `json:"redis"`
}

type SSHConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type ApiConfig struct {
	Token   string `json:"token"`
	BaseUrl string `json:"baseUrl"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       string `json:"db"`
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

func GetAPIConfig(filename string) (ApiConfig, error) {
	if jsonParm, err := loadConfig(filename); err != nil {
		return ApiConfig{}, err
	} else {
		return *jsonParm.ApiConfig, nil
	}
}

func GetRedisConfig(filename string) (RedisConfig, error) {
	if jsonParm, err := loadConfig(filename); err != nil {
		return RedisConfig{}, err
	} else {
		return *jsonParm.RedisConfig, nil
	}
}
