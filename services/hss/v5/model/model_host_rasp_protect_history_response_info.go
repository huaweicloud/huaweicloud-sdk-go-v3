package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostRaspProtectHistoryResponseInfo struct {

	// 服务器ip
	HostIp *string `json:"host_ip,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 告警时间
	AlarmTime *int64 `json:"alarm_time,omitempty"`

	// 威胁类型
	ThreatType *string `json:"threat_type,omitempty"`

	// 告警级别
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// 源IP
	SourceIp *string `json:"source_ip,omitempty"`

	// 攻击URL
	AttackedUrl *string `json:"attacked_url,omitempty"`
}

func (o HostRaspProtectHistoryResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostRaspProtectHistoryResponseInfo struct{}"
	}

	return strings.Join([]string{"HostRaspProtectHistoryResponseInfo", string(data)}, " ")
}
