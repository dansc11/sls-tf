package app

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/dansc11/sls-tf/app/types/terraform"
)

func Plan(workDir string) {
	slsConfig, err := loadServerlessConfig(workDir)

	if err != nil {
		log.Fatal(err)
	}

	writeSlsTfYml(workDir, slsConfig)

	// Replace with a tfvars file instead of string args
	tfVars := terraform.NewTerraformVariables(slsConfig)

	if err := runTerraformInit(workDir); err != nil {
		log.Fatal(err)
	}

	if err := runTerraformPlan(workDir, tfVars); err != nil {
		log.Fatal(err)
	}
}

func Deploy(workDir string) {
	slsConfig, err := loadServerlessConfig(workDir)

	if err != nil {
		log.Fatal(err)
	}

	writeSlsTfYml(workDir, slsConfig)

	// Replace with a tfvars file instead of string args
	tfVars := terraform.NewTerraformVariables(slsConfig)

	if err := runTerraformInit(workDir); err != nil {
		log.Fatal(err)
	}

	if err := runTerraformApply(workDir, tfVars); err != nil {
		log.Fatal(err)
	}
}

func runTerraformPlan(workDir string, variables terraform.TerraformVariables) error {
	command := fmt.Sprintf("terraform -chdir=%s plan %s", workDir, variables)

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

func runTerraformApply(workDir string, variables terraform.TerraformVariables) error {
	command := fmt.Sprintf("terraform -chdir=%s apply -auto-approve %s", workDir, variables)

	log.Printf("Preparing Terraform Apply command: %s", command)

	splitCommand := strings.Split(command, " ")

	cmd := exec.Command(splitCommand[0], splitCommand[1:]...)

	output, err := cmd.CombinedOutput()

	log.Println(string(output))

	if err != nil {
		return err
	}

	return nil
}

func runTerraformInit(workDir string) error {
	command := fmt.Sprintf("terraform -chdir=%s init", workDir)

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
