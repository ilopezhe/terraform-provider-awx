terraform {
  required_providers {
    awx = {
      source  = "github.com/ilopezhe/awx"
      version = "24.6.1"
    }
    time = {
      source = "hashicorp/time"
      version = "0.12.0"
    }
    random = {
      source = "hashicorp/random"
      version = "3.6.2"
    }
  }
}
