package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventSchemaVersionsRequest Request Object
type ListEventSchemaVersionsRequest struct {

	// 指定查询的事件模型ID
	SchemaId string `json:"schema_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，不能小于1或大于1000
	Limit *int32 `json:"limit,omitempty"`

	// 指定查询排序
	Sort *string `json:"sort,omitempty"`

	// 指定查询的事件模型版本号
	Version *string `json:"version,omitempty"`
}

func (o ListEventSchemaVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSchemaVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListEventSchemaVersionsRequest", string(data)}, " ")
}
