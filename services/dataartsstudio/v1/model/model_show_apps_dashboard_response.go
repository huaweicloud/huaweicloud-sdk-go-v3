package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppsDashboardResponse Response Object
type ShowAppsDashboardResponse struct {

	// 统计信息仪表板
	Dashboards     *[]StatisticForDashboard `json:"dashboards,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowAppsDashboardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppsDashboardResponse struct{}"
	}

	return strings.Join([]string{"ShowAppsDashboardResponse", string(data)}, " ")
}
