package hello

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type Hello struct {
	ID       int    `json:"id"`
	Greeting string `json:"greeting"`
	JSON     string `json:"json"`
}

func helloCols() []*plugin.Column {
	return []*plugin.Column{
		{Name: "id", Type: proto.ColumnType_INT, Description: "id"},
		{Name: "greeting", Type: proto.ColumnType_STRING, Description: "greeting"},
		{Name: "json", Type: proto.ColumnType_JSON, Description: "json"},
	}
}

func getGreeting(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	id := int(quals["id"].GetInt64Value())	
	plugin.Logger(ctx).Info("getGreeting", "number", id)
	greeting := Hello{id, "Hello", "{\"hello\": \"world\"}"}
	return &greeting, nil
}

func listGreeting(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 1; i <= 3; i++ {
		plugin.Logger(ctx).Info("listGreeting", "number", i)		
		greeting := Hello{i, "Hello", "{\"hello\": \"world\"}"}
		d.StreamListItem(ctx, &greeting)
	}
	return nil, nil
}
