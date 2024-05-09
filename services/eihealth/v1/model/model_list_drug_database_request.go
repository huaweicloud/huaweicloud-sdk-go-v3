package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDrugDatabaseRequest Request Object
type ListDrugDatabaseRequest struct {

	// 数据库名称搜索
	SearchKey *string `json:"search_key,omitempty"`

	// 数据库类型搜索
	Type *string `json:"type,omitempty"`

	// 排序规则 目前默认时间降序，支持根据create_time|update_time
	SortKey *string `json:"sort_key,omitempty"`

	// 排序规则 目前默认时间降序
	SortDir *string `json:"sort_dir,omitempty"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDrugDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDrugDatabaseRequest struct{}"
	}

	return strings.Join([]string{"ListDrugDatabaseRequest", string(data)}, " ")
}
