package wordpress

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-wordpress",
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},		
		TableMap: map[string]*plugin.Table{
			"hello_wordpress":         tableHelloWordPress(ctx),
			"wordpress_post":         tableWordPressPost(ctx),
		},
	}
	return p
}

