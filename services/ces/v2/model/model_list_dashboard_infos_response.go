package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDashboardInfosResponse Response Object
type ListDashboardInfosResponse struct {

	// **参数描述**： 监控看板列表
	Dashboards     *[]DashBoardInfo `json:"dashboards,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListDashboardInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDashboardInfosResponse struct{}"
	}

	return strings.Join([]string{"ListDashboardInfosResponse", string(data)}, " ")
}
