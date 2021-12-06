provider "aws" {
  region = var.region

  default_tags {
    tags = {
      Service     = var.service_name
      Environment = var.stage
    }
  }
}

module "service" {
  source = "github.com/dansc11/sls-tf/terraform"
}

variable "service_name" {
  type = string
}

variable "region" {
  type = string
}

variable "stage" {
  type = string
}
