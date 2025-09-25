package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostRaspProtectHistoryResponseInfo struct {

	// **参数解释**: 服务器IP **取值范围**: 字符长度0-64位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 告警时间，单位毫秒。 **取值范围**: 最小值0，最大值4070880000000
	AlarmTime *int64 `json:"alarm_time,omitempty"`

	// **参数解释**： 威胁类型 **取值范围**： 字符长度0-128位
	ThreatType *string `json:"threat_type,omitempty"`

	// **参数解释**: 告警等级 **取值范围**: - 1 : 紧急。 - 2 : 重要。 - 3 : 次要。 - 4 : 提示。
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// **参数解释**： 攻击源IP **取值范围**： 字符长度0-128位
	SourceIp *string `json:"source_ip,omitempty"`

	// **参数解释**： 攻击源URL **取值范围**： 字符长度0-2000位
	AttackedUrl *string `json:"attacked_url,omitempty"`
}

func (o HostRaspProtectHistoryResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostRaspProtectHistoryResponseInfo struct{}"
	}

	return strings.Join([]string{"HostRaspProtectHistoryResponseInfo", string(data)}, " ")
}
