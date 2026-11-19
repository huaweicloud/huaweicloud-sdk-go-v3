package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HoneypotEventResponseInfo struct {

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器的唯一标识ID **取值范围**: 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **取值范围**: 字符长度1-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 服务器私有IP **取值范围**: 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器弹性IP地址 **取值范围**: IPv4格式（长度7-15位）、IPv6格式（长度15-39位）
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 威胁等级 **取值范围**： - Security：安全 - Low: 低危 - Medium: 中危 - High: 高危 - Critical: 危急
	Severity *string `json:"severity,omitempty"`

	// **参数解释**: 攻击源IP **取值范围**: 字符长度1-128位
	AttackIp *string `json:"attack_ip,omitempty"`

	// **参数解释**: 攻击次数 **取值范围**: 最小值0，最大值2147483547
	AttackNum *int32 `json:"attack_num,omitempty"`

	// **参数解释**: 最近攻击时间，毫秒级时间戳 **取值范围**: 最小值0，最大值9223372036854775807
	AttackLastTime *int64 `json:"attack_last_time,omitempty"`

	// **参数解释**: 攻击开始时间，毫秒级时间戳 **取值范围**: 最小值0，最大值9223372036854775807
	AttackStartTime *int64 `json:"attack_start_time,omitempty"`

	// **参数解释**: 设备指纹 **取值范围**: 字符长度1-512位
	DeviceFingerprint *string `json:"device_fingerprint,omitempty"`
}

func (o HoneypotEventResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HoneypotEventResponseInfo struct{}"
	}

	return strings.Join([]string{"HoneypotEventResponseInfo", string(data)}, " ")
}
