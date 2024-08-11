package main

import (
	"github.com/judell/steampipe-plugin-wordpress/wordpress"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: wordpress.Plugin})
}
