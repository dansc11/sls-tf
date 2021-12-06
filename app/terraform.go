package app

import (
	"fmt"
	"log"
	"os/exec"
	"sls-tf/app/types/terraform"
	"strings"
)

func Plan(serverlessYmlPath string) {
	slsConfig, err := loadServerlessConfig(serverlessYmlPath)

	if err != nil {
		log.Fatal(err)
	}

	writeSlsTfYml(slsConfig)

	// Replace with a tfvars file instead of string args
	tfVars := terraform.NewTerraformVariables(slsConfig)

	if err := runTerraformInit(); err != nil {
		log.Fatal(err)
	}

	if err := runTerraformPlan(tfVars); err != nil {
		log.Fatal(err)
	}
}

func runTerraformPlan(variables terraform.TerraformVariables) error {
	command := fmt.Sprintf("terraform plan %s", variables)

	log.Printf("Preparing Terraform Plan command: %s", command)

	splitCommand := strings.Split(command, " ")

	cmd := exec.Command(splitCommand[0], splitCommand[1:]...)

	output, err := cmd.CombinedOutput()

	log.Println(string(output))

	if err != nil {
		return err
	}

	return nil
}

func runTerraformInit() error {
	command := "terraform init"

	log.Printf("Preparing Terraform Init command: %s", command)

	splitCommand := strings.Split(command, " ")

	cmd := exec.Command(splitCommand[0], splitCommand[1:]...)

	output, err := cmd.CombinedOutput()

	log.Println(string(output))

	if err != nil {
		return err
	}

	return nil
}
