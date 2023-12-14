package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListQueriesRequestBody struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	// 项目ID。
	ProjectId string `json:"project_id"`

	// 偏移量，表示从此偏移量开始查询，offset>=0。
	Offset string `json:"offset"`

	// 每页显示的条目数量。
	Limit string `json:"limit"`

	// 查询条件数组
	Conditions []ListQueriesCondition `json:"conditions"`

	// 排序字段
	OrderBy *string `json:"order_by,omitempty"`

	// 固定值db_queries
	Target string `json:"target"`
}

func (o ListQueriesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesRequestBody struct{}"
	}

	return strings.Join([]string{"ListQueriesRequestBody", string(data)}, " ")
}
