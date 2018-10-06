package main

import (
  "github.com/hashicorp/terraform/plugin"
  "github.com/extenda/fleet-manager-terraform-provider/fleetmanager"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: fleetmanager.Provider,
  })
}