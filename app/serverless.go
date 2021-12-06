package app

import (
	"io/ioutil"
	"sls-tf/app/types/serverless"

	"gopkg.in/yaml.v3"
)

func loadServerlessConfig(path string) (serverless.ServerlessConfig, error) {
	var config serverless.ServerlessConfig

	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return config, err
	}

	return config, nil
}

func writeSlsTfYml(config serverless.ServerlessConfig) error {
	data, err := yaml.Marshal(&config.Functions)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("./.sls-tf/functions.yml", data, 0777); err != nil {
		return err
	}

	return nil
}
