package terraform

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/dansc11/sls-tf/app/types/serverless"
)

type TerraformVariables struct {
	Region      string `tf:"region"`
	Stage       string `tf:"stage"`
	ServiceName string `tf:"service_name"`
}

func (v TerraformVariables) toTagMap() map[string]string {
	var tagMap map[string]string = make(map[string]string)

	variablesType := reflect.TypeOf(v)
	variablesInstance := reflect.ValueOf(v)

	for i := 0; i < variablesType.NumField(); i++ {
		tfVariableName := variablesType.Field(i).Tag.Get("tf")
		structPropertyName := variablesType.Field(i).Name

		variableValue := reflect.Indirect(variablesInstance).FieldByName(structPropertyName)

		tagMap[tfVariableName] = variableValue.String()
	}

	return tagMap
}

func (v TerraformVariables) toFilteredTagMap() map[string]string {
	filtered := make(map[string]string)

	tagMap := v.toTagMap()

	for tag, value := range tagMap {
		if value == "" {
			continue
		}

		filtered[tag] = value
	}

	return filtered
}

func (v TerraformVariables) String() string {
	var varsString string

	tagMap := v.toFilteredTagMap()

	// Sort the keys
	keys := make([]string, 0)

	for key, _ := range tagMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		varsString = fmt.Sprintf("%s -var %s=%s", varsString, key, tagMap[key])
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
