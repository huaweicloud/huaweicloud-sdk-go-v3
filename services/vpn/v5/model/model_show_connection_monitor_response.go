package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectionMonitorResponse Response Object
type ShowConnectionMonitorResponse struct {
	ConnectionMonitor *ConnectionMonitorInfo `json:"connection_monitor,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConnectionMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionMonitorResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectionMonitorResponse", string(data)}, " ")
}
