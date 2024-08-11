package wordpress

import (
	"context"
  "reflect"

	"github.com/sogko/go-wordpress"	
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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

type ListFunc func(context.Context, interface{}, int, int) (interface{}, *wordpress.Response, error)

func paginate(ctx context.Context, d *plugin.QueryData, listFunc ListFunc, options interface{}) error {
	perPage := 100 // Adjust this value based on your needs and API limits
	offset := 0

	for {
		items, _, err := listFunc(ctx, options, perPage, offset)
		if err != nil {
			plugin.Logger(ctx).Error("wordpress.paginate", "query_error", err)
			return err
		}

		itemsSlice := reflect.ValueOf(items)
		for i := 0; i < itemsSlice.Len(); i++ {
			d.StreamListItem(ctx, itemsSlice.Index(i).Interface())
		}

		// Check if we've reached the end of the items
		if itemsSlice.Len() < perPage {
			break
		}

		// Update the offset for the next page
		offset += perPage
	}

	return nil
}
