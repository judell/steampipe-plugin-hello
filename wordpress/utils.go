package wordpress

import (
	"context"
  //"time"

	"github.com/sogko/go-wordpress"	
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"	
)

func getDate(ctx context.Context, d *transform.TransformData) (interface{}, error) {
  post := d.Value.(*wordpress.Post)
	date := post.Date.Time
	return date, nil
}

func getTitle(ctx context.Context, d *transform.TransformData) (interface{}, error) {
  post := d.Value.(*wordpress.Post)
	title := post.Title.Rendered
	return title, nil
}


