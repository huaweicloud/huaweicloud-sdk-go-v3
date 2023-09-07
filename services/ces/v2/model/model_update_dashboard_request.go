package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDashboardRequest Request Object
type UpdateDashboardRequest struct {

	// 监控看板id，以db开头，包含22个字母和数字例：db16564943172807wjOmoLyn'
	DashboardId string `json:"dashboard_id"`

	Body *UpdateDashboardRequestBody `json:"body,omitempty"`
}

func (o UpdateDashboardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDashboardRequest struct{}"
	}

	return strings.Join([]string{"UpdateDashboardRequest", string(data)}, " ")
}
