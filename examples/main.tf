terraform {
  required_providers {
    jsonconf = {
      version = "~> 0.1"
      source  = "github.com/prividen/jsonconf"
    }
  }
}

provider "jsonconf" {}

data "jsonconf_nodes" "all" {}

output "datacenters" {
  value = {
    for k, v in data.jsonconf_nodes.all.nodes : k => v.dc
  }
}
