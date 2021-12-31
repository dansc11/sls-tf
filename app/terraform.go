package app

import (
	"log"

	"github.com/dansc11/sls-tf/app/terraform"
)

func Plan(workDir string) {
	slsConfig, err := loadServerlessConfig(workDir)

	if err != nil {
		log.Fatal(err)
	}

	writeSlsTfYml(workDir, slsConfig)

	// Replace with a tfvars file instead of string args
	tfVars := terraform.NewTerraformVariables(slsConfig)

	tfExecutor, err := terraform.NewExecutor(workDir)
	if err != nil {
		log.Fatal(err)
	}

	tfExecutor.SetVariables(tfVars)

	if err := tfExecutor.Init(); err != nil {
		log.Fatal(err)
	}

	if err := tfExecutor.Plan(); err != nil {
		log.Fatal(err)
	}
}

func Deploy(workDir string) {
	slsConfig, err := loadServerlessConfig(workDir)

	if err != nil {
		log.Fatal(err)
	}

	writeSlsTfYml(workDir, slsConfig)

	// TODO: Replace with a tfvars file instead of string args
	tfVars := terraform.NewTerraformVariables(slsConfig)

	tfExecutor, err := terraform.NewExecutor(workDir)
	if err != nil {
		log.Fatal(err)
	}

	tfExecutor.SetVariables(tfVars)

	if err := tfExecutor.Init(); err != nil {
		log.Fatal(err)
	}

	if err := tfExecutor.Apply(); err != nil {
		log.Fatal(err)
	}
}
