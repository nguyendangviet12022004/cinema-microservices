terraform {
  backend "s3" {
    bucket = "vietnd-terraform-state-bucket"
    key    = "dev/terraform.tfstate"
    region = "ap-southeast-1"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.33.0"
    }
  }
}

provider "aws" {
  region = "ap-southeast-1"

}
