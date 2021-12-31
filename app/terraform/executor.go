package terraform

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type operation int

const (
	Init operation = iota
	Plan
	Apply
	Destroy
)

func (o operation) String() string {
	return [...]string{
		"init",
		"plan",
		"apply -auto-approve",
		"destroy",
	}[o]
}

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
	return e.runCommand(Plan.String())
}

func (e *terraformExecutor) Apply() error {
	return e.runCommand(Apply.String())
}

func (e *terraformExecutor) Init() error {
	return e.runCommand(Init.String())
}

func (e *terraformExecutor) Destroy() error {
	return e.runCommand(Destroy.String())
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
