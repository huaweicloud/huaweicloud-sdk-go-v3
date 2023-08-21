package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDashboardRequest Request Object
type DeleteDashboardRequest struct {

	// 仪表盘id
	Id string `json:"id"`

	// 是否删除图表
	IsDeleteCharts bool `json:"is_delete_charts"`
}

func (o DeleteDashboardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDashboardRequest struct{}"
	}

	return strings.Join([]string{"DeleteDashboardRequest", string(data)}, " ")
}
