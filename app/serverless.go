package app

import (
	"io/ioutil"

	"github.com/dansc11/sls-tf/app/types/serverless"

	"gopkg.in/yaml.v3"
)

func loadServerlessConfig(workDir string) (serverless.ServerlessConfig, error) {
	var config serverless.ServerlessConfig

	yamlFile, err := ioutil.ReadFile(workDir + "/serverless.yml")

	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return config, err
	}

	return config, nil
}

func writeSlsTfYml(workDir string, config serverless.ServerlessConfig) error {
	data, err := yaml.Marshal(&config.Functions)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(workDir+"/.sls-tf/functions.yml", data, 0777); err != nil {
		return err
	}

	return nil
}
