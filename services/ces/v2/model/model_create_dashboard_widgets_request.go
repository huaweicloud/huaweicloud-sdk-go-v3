package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDashboardWidgetsRequest Request Object
type CreateDashboardWidgetsRequest struct {

	// 监控看板id，以db开头，包含22个字母和数字例：db16564943172807wjOmoLyn'
	DashboardId string `json:"dashboard_id"`

	Body *[]interface{} `json:"body,omitempty"`
}

func (o CreateDashboardWidgetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDashboardWidgetsRequest struct{}"
	}

	return strings.Join([]string{"CreateDashboardWidgetsRequest", string(data)}, " ")
}
