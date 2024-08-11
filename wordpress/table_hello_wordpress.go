package wordpress

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type HelloWordPress struct {
	Greeting string `json:"greeting"`
}

func tableHelloWordPress(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hello_wordpress",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listGreeting,
		},
		Columns: helloCols(),
	}
}

func helloCols() []*plugin.Column {
	return []*plugin.Column{
		{Name: "greeting", Type: proto.ColumnType_STRING, Description: "greeting"},
	}
}

func listGreeting(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	greeting := HelloWordPress{"Hello"}
	d.StreamListItem(ctx, &greeting)
	return nil, nil
}
