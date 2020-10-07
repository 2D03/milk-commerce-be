package conf

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

var Config *config

type config struct {
	MainDBName string
	LogDBName  string
}

func init() {
	env := os.Getenv("env")
	switch env {
	case "local":
		Config = &config{
			MainDBName: "milk_commerce_local",
			LogDBName:  "milk_commerce_local_log",
		}
	case "dev":
		Config = &config{
			MainDBName: "milk_commerce_dev",
			LogDBName:  "milk_commerce_dev_log",
		}
	case "prd":
		Config = &config{
			MainDBName: "milk_commerce_prd",
			LogDBName:  "milk_commerce_prd_log",
		}
	}
}

func GetConfigDB() (map[string]string, error) {
	var configMap map[string]string
	configStr := os.Getenv("config")
	decoded, err := base64.URLEncoding.DecodeString(configStr)
	if err != nil {
		fmt.Println("[Parse config] Convert B64 config string error: " + err.Error())
		return nil, err
	}
	err = json.Unmarshal(decoded, &configMap)
	if err != nil {
		fmt.Println("[Parse config] Parse JSON with config string error: " + err.Error())
		return nil, err
	}

	return configMap, err
}
