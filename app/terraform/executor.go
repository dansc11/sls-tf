package terraform

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type terraformExecutor struct {
	WorkDir             string
	variables           TerraformVariables
	TerraformBinaryPath string
}

func NewExecutor(workDir string) (terraformExecutor, error) {
	var executor terraformExecutor

	executor.WorkDir = workDir
	executor.TerraformBinaryPath = "terraform"

	return executor, nil
}

func (e *terraformExecutor) SetVariables(variables TerraformVariables) {
	e.variables = variables
}

func (e *terraformExecutor) buildCommand(operation string) string {
	return fmt.Sprintf(
		"%s -chdir=%s %s %s",
		e.TerraformBinaryPath,
		e.WorkDir,
		operation,
		e.variables,
	)
}

func (e *terraformExecutor) Plan() error {
	// TODO: Replace with enum
	return e.runCommand("plan")
}

func (e *terraformExecutor) Apply() error {
	// TODO: pass args separately
	return e.runCommand("apply -auto-approve")
}

func (e *terraformExecutor) Init() error {
	return e.runCommand("init")
}

func (e *terraformExecutor) runCommand(operation string) error {
	command := e.buildCommand(operation)

	log.Printf("Preparing Terraform command: %s", command)

	splitCommand := strings.Split(command, " ")

	cmd := exec.Command(splitCommand[0], splitCommand[1:]...)

	output, err := cmd.CombinedOutput()

	log.Println(string(output))

	if err != nil {
		return err
	}

	return nil
}
