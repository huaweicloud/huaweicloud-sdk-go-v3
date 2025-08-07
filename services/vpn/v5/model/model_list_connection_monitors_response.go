package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionMonitorsResponse Response Object
type ListConnectionMonitorsResponse struct {

	// VPN连接监控列表
	ConnectionMonitors *[]ConnectionMonitorInfo `json:"connection_monitors,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListConnectionMonitorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionMonitorsResponse struct{}"
	}

	return strings.Join([]string{"ListConnectionMonitorsResponse", string(data)}, " ")
}
