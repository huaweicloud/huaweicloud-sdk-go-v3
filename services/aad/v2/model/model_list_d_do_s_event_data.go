package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListDDoSEventData struct {

	// 防护IP
	ZoneIp *string `json:"zone_ip,omitempty"`

	// 开始时间（毫秒时间戳）
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间（毫秒时间戳）
	EndTime *string `json:"end_time,omitempty"`

	// 攻击流量峰值，单位“kbps”
	MaxDropKbps *string `json:"max_drop_kbps,omitempty"`

	// 攻击报文数峰值，单位“pps”
	MaxDropPps *string `json:"max_drop_pps,omitempty"`

	// 入流量峰值，单位“kbps”
	MaxInKbps *string `json:"max_in_kbps,omitempty"`

	// 入报文数峰值，单位“pps”
	MaxInPps *string `json:"max_in_pps,omitempty"`

	// 攻击类型
	AttackTypes *string `json:"attack_types,omitempty"`

	// 攻击源IP
	AttackIps *string `json:"attack_ips,omitempty"`

	// 攻击IP描述
	AttackIpsDesc *string `json:"attack_ips_desc,omitempty"`

	// 攻击状态
	AttackStatus *string `json:"attack_status,omitempty"`
}

func (o ListDDoSEventData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDoSEventData struct{}"
	}

	return strings.Join([]string{"ListDDoSEventData", string(data)}, " ")
}
