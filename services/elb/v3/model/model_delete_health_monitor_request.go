package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteHealthMonitorRequest struct {

	// 健康检查ID。
	HealthmonitorId string `json:"healthmonitor_id" xml:"healthmonitor_id"`
}

func (o DeleteHealthMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"DeleteHealthMonitorRequest", string(data)}, " ")
}
