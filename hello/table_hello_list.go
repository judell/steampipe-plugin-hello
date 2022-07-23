package hello

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
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

func listGreeting(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 1; i <= 3; i++ {
		plugin.Logger(ctx).Info("listGreeting", "number", i)		
		greeting := Hello{i, "Hello", "{\"hello\": \"world\"}"}
		d.StreamListItem(ctx, &greeting)
	}
	return nil, nil
}
