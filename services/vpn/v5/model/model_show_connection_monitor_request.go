package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectionMonitorRequest Request Object
type ShowConnectionMonitorRequest struct {

	// VPN连接监控的ID
	ConnectionMonitorId string `json:"connection_monitor_id"`
}

func (o ShowConnectionMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionMonitorRequest struct{}"
	}

	return strings.Join([]string{"ShowConnectionMonitorRequest", string(data)}, " ")
}
