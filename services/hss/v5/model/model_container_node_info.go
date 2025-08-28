package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerNodeInfo 容器节点列表
type ContainerNodeInfo struct {

	// **参数解释**: Agent ID **取值范围**: 字符长度0-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 服务器ID **取值范围**: 字符长度0-128位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 节点名称 **取值范围**: 字符长度0-128位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器状态 **取值范围**: 包含如下4种。   - ACTIVE：正在运行。   - SHUTOFF：关机。   - BUILDING：创建中。   - ERROR：故障。
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**: Agent状态 **取值范围**: 包含如下3种。   - not_installed：未安装。   - online：在线。   - offline：离线。
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: 防护状态 **取值范围**: 包含如下2种。   - closed ：关闭。   - opened ：开启。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**: 防护是否中断 **取值范围**:   - true：防护中断。   - false：防护未中断。
	ProtectInterrupt *bool `json:"protect_interrupt,omitempty"`

	// **参数解释**: 防护是否降级 **取值范围**:   - true：防护降级。   - false：防护未降级。
	ProtectDegradation *bool `json:"protect_degradation,omitempty"`

	// **参数解释**: 防护降级原因 **取值范围**: 字符长度1-32位
	DegradationReason *string `json:"degradation_reason,omitempty"`

	// **参数解释**: 用来识别cce容器节点和自建节点的标签 **取值范围**: 包含如下3种。 - cce：cce节点 - self：自建节点 - other：其他节点
	ContainerTags *string `json:"container_tags,omitempty"`

	// **参数解释**: 私有IP地址 **取值范围**: 字符长度0-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 弹性公网IP地址 **取值范围**: 字符长度0-128位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 主机安全配额ID（UUID） **取值范围**: 字符长度0-128位
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**: 服务器组名称 **取值范围**: 字符长度1-128位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 所属企业项目名称 **取值范围**: 字符长度0-256位
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// **参数解释**: 云主机安全检测结果 **取值范围**: 包含如下4种。 - undetected：未检测。 - clean：无风险。 - risk：有风险。 - scanning：检测中。
	DetectResult *string `json:"detect_result,omitempty"`

	// **参数解释**: 资产风险 **取值范围**: 取值0-2097152
	Asset *int32 `json:"asset,omitempty"`

	// **参数解释**: 漏洞风险 **取值范围**: 取值0-2097152
	Vulnerability *int32 `json:"vulnerability,omitempty"`

	// **参数解释**: 入侵风险 **取值范围**: 取值0-2097152
	Intrusion *int32 `json:"intrusion,omitempty"`

	// **参数解释**: 策略组ID **取值范围**: 字符长度1-128位
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// **参数解释**: 策略组名称 **取值范围**: 字符长度1-128位
	PolicyGroupName *string `json:"policy_group_name,omitempty"`
}

func (o ContainerNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerNodeInfo struct{}"
	}

	return strings.Join([]string{"ContainerNodeInfo", string(data)}, " ")
}
