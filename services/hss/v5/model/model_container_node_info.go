package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerNodeInfo 容器节点列表
type ContainerNodeInfo struct {

	// **参数解释** AgentId标识 **取值范围** 只能由英文字母、数字、特殊字符组成, 长度范围为[0-64]个字符
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释** 服务器ID **取值范围** 只能由英文字母、数字、特殊字符组成, 长度范围为[0-128]个字符
	HostId *string `json:"host_id,omitempty"`

	// **参数解释** 服务器名称 **取值范围** 只能由中文字符、英文字母、数字、特殊字符组成, 长度范围为[0-128]个字符
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** 服务器状态 **取值范围** - ACTIVE ：正在运行 - SHUTOFF ：关机 - BUILDING ：创建中 - ERROR ：故障
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释** Agent状态 **取值范围**         - not_installed ：未安装 - online ：在线 - offline ：离线
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释** 防护状态 **取值范围**         - closed ：防护关闭状态 - opened ：防护开启状态
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释** 防护是否中断 **取值范围**         - true: 防护中断 - false: 防护未中断
	ProtectInterrupt *bool `json:"protect_interrupt,omitempty"`

	// **参数解释** 防护是否降级 **取值范围**         - true: 防护降级 - false: 防护未降级
	ProtectDegradation *bool `json:"protect_degradation,omitempty"`

	// **参数解释** 防护降级原因 **取值范围**         只能由中文字符、英文字母、数字、特殊字符组成, 长度范围为[1-32]个字符
	DegradationReason *string `json:"degradation_reason,omitempty"`

	// **参数解释** 标签：用来识别cce集群节点和自建集群节点 **取值范围**         - cce：cce节点 - self：自建节点 - other：其它节点
	ContainerTags *string `json:"container_tags,omitempty"`

	// **参数解释** 私有IP地址 **取值范围**         只能由数字、特殊字符组成, 长度范围为[0-128]个字符
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释** 弹性公网IP地址 **取值范围**         只能由数字、特殊字符组成, 长度范围为[0-128]个字符
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释** 主机安全配额ID（UUID） **取值范围**         只能由中文字符、英文字母、数字、特殊字符组成, 长度范围为[0-128]个字符
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释** 服务器组名称 **取值范围**         只能由中文字符、英文字母、数字、特殊字符组成, 长度范围为[0-128]个字符
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释** 所属企业项目名称 **取值范围**         只能由中文字符、英文字母、数字、特殊字符组成, 长度范围为[0-256]个字符
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// **参数解释** 云主机安全检测结果 **取值范围**           - undetected ：未检测 - clean ：无风险 - risk ：有风险 - scanning ：检测中
	DetectResult *string `json:"detect_result,omitempty"`

	// **参数解释** 资产风险 **取值范围**         0-2147483647
	Asset *int32 `json:"asset,omitempty"`

	// **参数解释** 漏洞风险 **取值范围**         0-2147483647
	Vulnerability *int32 `json:"vulnerability,omitempty"`

	// **参数解释** 入侵风险 **取值范围**         0-2147483647
	Intrusion *int32 `json:"intrusion,omitempty"`

	// **参数解释** 策略组ID **取值范围**         只能由英文字母、数字、特殊字符组成, 长度范围为[1-128]个字符
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// **参数解释** 策略组名称 **取值范围**         只能由中文字母、英文字母、数字、特殊字符组成, 长度范围为[1-128]个字符
	PolicyGroupName *string `json:"policy_group_name,omitempty"`
}

func (o ContainerNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerNodeInfo struct{}"
	}

	return strings.Join([]string{"ContainerNodeInfo", string(data)}, " ")
}
