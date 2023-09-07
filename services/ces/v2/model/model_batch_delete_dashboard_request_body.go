package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteDashboardRequestBody struct {

	// 监控看板id列表
	DashboardIds *[]string `json:"dashboard_ids,omitempty"`
}

func (o BatchDeleteDashboardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDashboardRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteDashboardRequestBody", string(data)}, " ")
}
