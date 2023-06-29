package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeFlowsResponse Response Object
type ListEdgeFlowsResponse struct {

	// **参数说明**：条件查询返回的总条数。
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：车辆流量，平均速度等统计信息列表
	Statistics     *[]OpenV2XStatisticsBody `json:"statistics,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListEdgeFlowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeFlowsResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeFlowsResponse", string(data)}, " ")
}
