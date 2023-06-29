package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApisDashboardResponse Response Object
type ShowApisDashboardResponse struct {

	// 统计信息仪表板
	Dashboards     *[]StatisticForDashboard `json:"dashboards,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowApisDashboardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApisDashboardResponse struct{}"
	}

	return strings.Join([]string{"ShowApisDashboardResponse", string(data)}, " ")
}
