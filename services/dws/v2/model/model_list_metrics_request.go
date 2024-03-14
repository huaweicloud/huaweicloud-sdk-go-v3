package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsRequest Request Object
type ListMetricsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 偏移量，表示从此偏移量开始查询，offset>=0。
	Offset int32 `json:"offset"`

	// 每页显示的条目数量，最大1000。
	Limit int32 `json:"limit"`

	// 排序字段。固定取值。 create_time：创建时间。
	OrderBy *string `json:"order_by,omitempty"`

	// 正序还是倒叙。固定取值。 asc：正序。 desc：倒序。
	SortBy *string `json:"sort_by,omitempty"`
}

func (o ListMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListMetricsRequest", string(data)}, " ")
}
