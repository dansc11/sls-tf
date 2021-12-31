package terraform

import (
	"testing"

	"github.com/dansc11/sls-tf/app/types/serverless"
	"github.com/stretchr/testify/assert"
)

func TestToTagMap(t *testing.T) {
	service := "api"
	region := "eu-west-2"
	stage := "dev"

	variables := TerraformVariables{
		Region:      region,
		Stage:       stage,
		ServiceName: service,
	}

	tagMap := variables.toTagMap()

	assert.Equal(t, region, tagMap["region"])
	assert.Equal(t, stage, tagMap["stage"])
	assert.Equal(t, service, tagMap["service_name"])
}

func TestToTagMapWithBlankValues(t *testing.T) {
	region := "eu-west-2"
	stage := "dev"

	variables := TerraformVariables{
		Region: region,
		Stage:  stage,
	}

	tagMap := variables.toTagMap()

	assert.Equal(t, region, tagMap["region"])
	assert.Equal(t, stage, tagMap["stage"])
	assert.Equal(t, "", tagMap["service_name"])
}

func TestToFilteredTagMap(t *testing.T) {
	service := "api"
	region := "eu-west-2"
	stage := "dev"

	variables := TerraformVariables{
		Region:      region,
		Stage:       stage,
		ServiceName: service,
	}

	tagMap := variables.toFilteredTagMap()

	assert.Equal(t, region, tagMap["region"])
	assert.Equal(t, stage, tagMap["stage"])
	assert.Equal(t, service, tagMap["service_name"])
}

func TestToFilteredTagMapWithBlankValues(t *testing.T) {
	region := "eu-west-2"
	stage := "dev"

	variables := TerraformVariables{
		Region: region,
		Stage:  stage,
	}

	tagMap := variables.toFilteredTagMap()

	assert.Equal(t, region, tagMap["region"])
	assert.Equal(t, stage, tagMap["stage"])
	assert.NotContains(t, tagMap, "service_name")
}

func TestString(t *testing.T) {
	service := "api"
	region := "eu-west-2"
	stage := "dev"

	variables := TerraformVariables{
		Region:      region,
		Stage:       stage,
		ServiceName: service,
	}

	want := "-var region=eu-west-2 -var service_name=api -var stage=dev"
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
