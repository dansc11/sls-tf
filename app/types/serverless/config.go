package serverless

type ServerlessConfig struct {
	Service   string                     `yaml:"service"`
	Provider  ServerlessConfigProvider   `yaml:"provider"`
	Functions []ServerlessConfigFunction `yaml:"functions"`
}

type ServerlessConfigProvider struct {
	Region string `yaml:"region"`
	Stage  string `yaml:"stage"`
}

type ServerlessConfigFunction struct {
	Name     string `yaml:"name"`
	Filename string `yaml:"filename"`
	Handler  string `yaml:"handler"`
	Runtime  string `yaml:"runtime"`
	Role     string `yaml:"role"`
}
