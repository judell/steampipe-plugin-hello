package hello

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
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