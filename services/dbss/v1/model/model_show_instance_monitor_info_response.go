package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceMonitorInfoResponse Response Object
type ShowInstanceMonitorInfoResponse struct {
	DiskInfos *DiskInfo `json:"disk_infos,omitempty"`

	// 系统信息
	SystemInfos *[]SystemInfo `json:"system_infos,omitempty"`

	// 流量统计
	TrafficInfos   *[]TrafficInfo `json:"traffic_infos,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowInstanceMonitorInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMonitorInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceMonitorInfoResponse", string(data)}, " ")
}
