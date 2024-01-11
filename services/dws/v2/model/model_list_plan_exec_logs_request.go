package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlanExecLogsRequest Request Object
type ListPlanExecLogsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 计划ID
	PlanId string `json:"plan_id"`

	// 查询条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListPlanExecLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlanExecLogsRequest struct{}"
	}

	return strings.Join([]string{"ListPlanExecLogsRequest", string(data)}, " ")
}
