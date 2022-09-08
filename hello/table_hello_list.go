package hello

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableHelloList(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hello_list",
		Description: "hello using List/ListConfig",
		List: &plugin.ListConfig{
			Hydrate: listGreeting,
		},
		Columns: helloCols(),
	}
}

