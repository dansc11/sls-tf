package terraform

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dansc11/sls-tf/app/types/serverless"
)

type TerraformVariables struct {
	Region      string `tf:"region"`
	Stage       string `tf:"stage"`
	ServiceName string `tf:"service_name"`
}

func (v TerraformVariables) String() string {
	var varsString string

	variablesType := reflect.TypeOf(v)
	variablesInstance := reflect.ValueOf(v)

	for i := 0; i < variablesType.NumField(); i++ {
		tfVariableName := variablesType.Field(i).Tag.Get("tf")
		structPropertyName := variablesType.Field(i).Name

		variableValue := reflect.Indirect(variablesInstance).FieldByName(structPropertyName)

		if variableValue.String() == "" {
			continue
		}

		varsString = fmt.Sprintf("%s -var %s=%s", varsString, tfVariableName, variableValue)
	}

	return strings.TrimPrefix(varsString, " ")
}

func NewTerraformVariables(slsConfig serverless.ServerlessConfig) TerraformVariables {
	return TerraformVariables{
		Region:      slsConfig.Provider.Region,
		Stage:       slsConfig.Provider.Stage,
		ServiceName: slsConfig.Service,
	}
}
