terraform {
  required_providers {
    logdemobuilder = {
      version = "0.2"
      source  = "daitest.leibcorp.com/util/logdemobuilder"
    }
  }
}

provider "logdemobuilder" {}

variable "log" {
  type    = string
  default = ""
}
variable "instance" {
  type    = string
  default = ""
}

data "logdemobuilder" "psl" {
  log = "my test log"
  instance = "my instance"
}
