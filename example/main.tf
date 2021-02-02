terraform {
  required_providers {
    logship = {
      version = "0.2"
      source  = "daidemos.com/util/logship"
    }
  }
}

provider "logship" {}

variable "log" {
  type    = string
  default = ""
}
variable "instance" {
  type    = string
  default = ""
}

data "logship" "psl" {
  log = "my test log"
  instance = "my instance"
}
