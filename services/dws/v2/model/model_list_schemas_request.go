package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchemasRequest Request Object
type ListSchemasRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 数据库名称
	DatabaseName string `json:"database_name"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// 排序字段
	SortDir *string `json:"sort_dir,omitempty"`

	// 查询关键词
	Keywords *string `json:"keywords,omitempty"`

	// 查询条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemasRequest struct{}"
	}

	return strings.Join([]string{"ListSchemasRequest", string(data)}, " ")
}
