package wordpress

import (
	"context"

	"github.com/sogko/go-wordpress"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableWordPressPost(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "wordpress_post",
		Description: "Represents a post in WordPress.",
		List: &plugin.ListConfig{
			Hydrate: listPosts,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "The post ID."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromValue().Transform(getTitle), Description: "The post title."},
			{Name: "content", Type: proto.ColumnType_JSON, Description: "The post content."},
			{Name: "author", Type: proto.ColumnType_INT, Description: "The post author ID."},
			{Name: "date", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromValue().Transform(getDate), Description: "The post publication date."},
		  {Name: "raw", Type: proto.ColumnType_JSON, Transform: transform.FromValue()},
		},
	}
}

func listPosts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	options := &wordpress.PostListOptions{}

	err = paginate(ctx, d, func(ctx context.Context, opts interface{}, perPage, offset int) (interface{}, *wordpress.Response, error) {
		options := opts.(*wordpress.PostListOptions)
		options.ListOptions.PerPage = perPage
		options.ListOptions.Offset = offset
		return conn.Posts.List(ctx, options)
	}, options)

	return nil, err
}

