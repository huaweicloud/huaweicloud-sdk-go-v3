package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHealthMonitorRequest Request Object
type ShowHealthMonitorRequest struct {

	// 健康检查ID。
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o ShowHealthMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthMonitorRequest", string(data)}, " ")
}
