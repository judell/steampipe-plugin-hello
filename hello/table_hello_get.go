package hello

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
)

func tableHelloGet(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hello_get",
		Description: "hello using Get/GetConfig",
		Get: &plugin.GetConfig{
			Hydrate: getGreeting,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "an int"},
			{Name: "greeting", Type: proto.ColumnType_STRING, Description: "a string"},
			{Name: "json", Type: proto.ColumnType_JSON, Description: "a json object"},
		},
	}
}

func getGreeting(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	    quals := d.KeyColumnQuals
	    id := int(quals["id"].GetInt64Value())	
		plugin.Logger(ctx).Info("listGreeting", "number", id)
		greeting := Hello{id, "Hello", "{\"hello\": \"world\"}"}
		return &greeting, nil
}
