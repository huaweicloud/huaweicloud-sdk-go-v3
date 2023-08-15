package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApiDashboardResponse Response Object
type ShowApiDashboardResponse struct {

	// 统计信息仪表板
	Dashboards     *[]StatisticForDashboard `json:"dashboards,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowApiDashboardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiDashboardResponse struct{}"
	}

	return strings.Join([]string{"ShowApiDashboardResponse", string(data)}, " ")
}
