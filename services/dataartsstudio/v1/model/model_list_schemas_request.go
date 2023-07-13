package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchemasRequest Request Object
type ListSchemasRequest struct {

	// 数据所在空间的id值
	Workspace string `json:"workspace"`

	// 数据连接id
	ConnectionId string `json:"connection_id"`

	// 数据库名称
	DatabaseName string `json:"database_name"`

	// 数据条数限制
	Limit *string `json:"limit,omitempty"`

	// 偏移量
	Offset *string `json:"offset,omitempty"`
}

func (o ListSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemasRequest struct{}"
	}

	return strings.Join([]string{"ListSchemasRequest", string(data)}, " ")
}
