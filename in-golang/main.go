package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Config struct {
	Description   string `yaml:"description"`
	Env           string `yaml:"env"`
	Version       string `yaml:"version"`
	ComponentName string `yaml:"componentName"`
}

func isEnvExist(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}

func main() {
	if !isEnvExist("NORBIX_CONFIG_FILE") {
		fmt.Println("No config file, exiting...")
		os.Exit(1)
	}

	cf := os.Getenv("NORBIX_CONFIG_FILE")
	filename, _ := filepath.Abs(cf)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	configMap := map[string]string{
		"description":   config.Description,
		"env":           config.Env,
		"version":       config.Version,
		"componentName": config.ComponentName,
	}

	//HINT: Contract
	configMapFromEnv := map[string]string{}

	if isEnvExist("NORBIX_ENV") {
		configMapFromEnv["env"] = os.ExpandEnv("${NORBIX_ENV}")
	}

	if isEnvExist("NORBIX_SEMVER") {
		configMapFromEnv["version"] = os.ExpandEnv("${NORBIX_SEMVER}")
	}

	//configMapFromEnv := map[string]string{
	//	"env":     os.ExpandEnv("${NORBIX_ENV}"),
	//	"version": os.ExpandEnv("${NORBIX_SEMVER}"),
	//}

	//fmt.Println(configMap)
	//fmt.Println(configMapFromEnv)
	//fmt.Println(config)

	//HINT: union maps
	for k, v := range configMapFromEnv {
		configMap[k] = v
	}

	//fmt.Println("==")
	//fmt.Println(configMap)
	//fmt.Println(configMapFromEnv)

	err = mapstructure.Decode(configMap, &config)
	if err != nil {
		panic(err)
	}
	//fmt.Println(config)

	//NOTE: hack for daemon thread
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("INFO - Running dragons with artEfact", config.Version, "...")
	}
}
