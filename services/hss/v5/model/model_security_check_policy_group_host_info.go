package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckPolicyGroupHostInfo 策略组关联的主机信息
type SecurityCheckPolicyGroupHostInfo struct {

	// **参数解释**: agent id **取值范围**: 字符长度1-64位，支持字母、数字、连字符、下划线，不能以特殊字符开头或结尾
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 主机ID **取值范围**: 字符长度1-128
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-128位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： agent版本 **取值范围**: 字符长度1-64位
	AgentVersion *string `json:"agent_version,omitempty"`

	// **参数解释**： 防护版本 **取值范围**: 字符长度1-256位
	Version *string `json:"version,omitempty"`

	// **参数解释**： 服务器组名称 **取值范围**: 字符长度1-256位
	GroupName *string `json:"group_name,omitempty"`
}

func (o SecurityCheckPolicyGroupHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckPolicyGroupHostInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckPolicyGroupHostInfo", string(data)}, " ")
}
