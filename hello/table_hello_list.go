package hello

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
)

func tableHelloList(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hello_list",
		Description: "hello using List/ListConfig",
		List: &plugin.ListConfig{
			Hydrate: listGreeting,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "an int"},
			{Name: "greeting", Type: proto.ColumnType_STRING, Description: "a string"},
			{Name: "json", Type: proto.ColumnType_JSON, Description: "a json object"},
		},
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
