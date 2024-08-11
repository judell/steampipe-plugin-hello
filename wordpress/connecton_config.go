package wordpress

import (
	"context"
	"fmt"

	"github.com/sogko/go-wordpress"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"

)


var ConfigSchema = map[string]*schema.Attribute{
	"endpoint": {
		Type: schema.TypeString,
	},
	"username": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
}

type PluginConfig struct {
	Endpoint    *string `cty:"endpoint"`
	Username     *string `cty:"username"`
	Password     *string `cty:"password"`
	}


func ConfigInstance() interface{} {
	return &PluginConfig{}
}


func connect(ctx context.Context, d *plugin.QueryData) (*wordpress.Client, error) {
	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "wordpress"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*wordpress.Client), nil
	}

	// Retrieve the connection config
	wordpressConfig := GetConfig(d.Connection)

	if *wordpressConfig.Endpoint == "" {
		return nil, fmt.Errorf("endpoint must be configured")
	}

	// Create auth transport
	tp := wordpress.BasicAuthTransport{
		Username: *wordpressConfig.Username,
		Password: *wordpressConfig.Password,
	}

	// Create client
	client, err := wordpress.NewClient(*wordpressConfig.Endpoint, tp.Client())
	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}

func GetConfig(connection *plugin.Connection) PluginConfig {
	if connection == nil || connection.Config == nil {
		return PluginConfig{}
	}

	config, _ := connection.Config.(PluginConfig)

	return config
}


