package hello

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-hello",
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"hello_get":         tableHelloGet(ctx),
			"hello_list":        tableHelloList(ctx),
			"hello_get_or_list": tableHelloGetOrList(ctx),
		},
	}
	return p
}
