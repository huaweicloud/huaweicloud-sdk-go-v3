package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConnectionMonitorRequest Request Object
type DeleteConnectionMonitorRequest struct {

	// VPN连接监控的ID
	ConnectionMonitorId string `json:"connection_monitor_id"`
}

func (o DeleteConnectionMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnectionMonitorRequest struct{}"
	}

	return strings.Join([]string{"DeleteConnectionMonitorRequest", string(data)}, " ")
}
