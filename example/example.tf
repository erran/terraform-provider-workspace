terraform {
  required_providers {
    workspace = {
      source = "errancarey.com/erran/workspace"
      version = ">= 0.0.1"
    }
  }
}

#provider "workspace" {
  #not_workspace = "not-default"
#}

resource "workspace_requirement" "x" {}
