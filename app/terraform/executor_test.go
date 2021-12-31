package terraform

import (
	"reflect"
	"testing"
)

func Test_operation_String(t *testing.T) {
	tests := []struct {
		name string
		o    operation
		want string
	}{
		{
			name: "test init command",
			o:    Init,
			want: "init",
		},
		{
			name: "test plan command",
			o:    Plan,
			want: "plan",
		},
		{
			name: "test apply command",
			o:    Apply,
			want: "apply -auto-approve",
		},
		{
			name: "test destroy command",
			o:    Destroy,
			want: "destroy -auto-approve",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("operation.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewExecutor(t *testing.T) {
	type args struct {
		workDir string
	}
	tests := []struct {
		name string
		args args
		want terraformExecutor
	}{
		{
			name: "test successful creation with workdir",
			args: args{
				workDir: "here",
			},
			want: terraformExecutor{
				WorkDir:             "here",
				TerraformBinaryPath: "terraform",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewExecutor(tt.args.workDir)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExecutor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_terraformExecutor_SetVariables(t *testing.T) {
	type fields struct {
		WorkDir             string
		variables           TerraformVariables
		TerraformBinaryPath string
	}
	type args struct {
		variables TerraformVariables
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test set variables",
			fields: fields{},
			args: args{
				variables: TerraformVariables{
					Region:      "eu-west-2",
					Stage:       "dev",
					ServiceName: "api",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &terraformExecutor{
				WorkDir:             tt.fields.WorkDir,
				variables:           tt.fields.variables,
				TerraformBinaryPath: tt.fields.TerraformBinaryPath,
			}
			e.SetVariables(tt.args.variables)
			if !reflect.DeepEqual(e.variables, tt.args.variables) {
				t.Errorf("TF variables = %v, want %v", e.variables, tt.args.variables)
			}
		})
	}
}
