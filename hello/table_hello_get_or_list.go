package hello

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableHelloGetOrList(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hello_get_or_list",
		Description: "hello using Get/GetConfig or List/ListConfig",
		Get: &plugin.GetConfig{
			Hydrate: getGreeting,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listGreeting,
		},
		Columns: helloCols(),
	}
}

