package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteDashboardsRequestBody struct {

	// **参数解释** 监控看板id列表 **约束限制** 包含的监控看板id对象个数为[1,30]
	DashboardIds *[]string `json:"dashboard_ids,omitempty"`
}

func (o BatchDeleteDashboardsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDashboardsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteDashboardsRequestBody", string(data)}, " ")
}
