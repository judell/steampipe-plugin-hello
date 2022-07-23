package hello

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableHelloGet(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hello_get",
		Description: "hello using Get/GetConfig",
		Get: &plugin.GetConfig{
			Hydrate: getGreeting,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: helloCols(),
	}
}

func getGreeting(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	if h.Item != nil {
		plugin.Logger(ctx).Info("getGreeting", "h.Item", h.Item)
	}
    quals := d.KeyColumnQuals
    id := int(quals["id"].GetInt64Value())	
	plugin.Logger(ctx).Info("getGreeting", "number", id)
	greeting := Hello{id, "Hello", "{\"hello\": \"world\"}"}
	return &greeting, nil
}
