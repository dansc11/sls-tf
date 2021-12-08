package terraform

import (
	"testing"

	"github.com/dansc11/sls-tf/app/types/serverless"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	variables := TerraformVariables{
		Region:      "eu-west-2",
		Stage:       "dev",
		ServiceName: "api",
	}

	want := "-var region=eu-west-2 -var stage=dev -var service_name=api"
	got := variables.String()

	assert.Equal(t, want, got)
}

func TestStringWithMissingValues(t *testing.T) {
	variables := TerraformVariables{
		Region: "eu-west-2",
		Stage:  "dev",
	}

	want := "-var region=eu-west-2 -var stage=dev"
	got := variables.String()

	assert.Equal(t, want, got)
}

func TestNewTerraformVariables(t *testing.T) {
	service := "api"
	region := "eu-west-2"
	stage := "dev"

	sls := serverless.ServerlessConfig{
		Service: service,
		Provider: serverless.ServerlessConfigProvider{
			Region: region,
			Stage:  stage,
		},
	}

	tfVars := NewTerraformVariables(sls)

	assert.Equal(t, service, tfVars.ServiceName)
	assert.Equal(t, region, tfVars.Region)
	assert.Equal(t, stage, tfVars.Stage)
}
