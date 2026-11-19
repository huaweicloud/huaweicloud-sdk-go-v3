package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerlessInstanceInfo 实例信息
type ServerlessInstanceInfo struct {

	// **参数解释**: 所属企业项目名称 **取值范围**: 字符长度0-65535位
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// **参数解释**: 实例名称 **取值范围**: 字符长度0-65535位
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**: 实例ID **取值范围**: 字符长度0-65535位
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 实例IP **取值范围**: 字符长度0-65535位
	InstanceIp *string `json:"instance_ip,omitempty"`

	// **参数解释**: 实例状态 **取值范围**: 字符长度0-65535位
	InstanceStatus *string `json:"instance_status,omitempty"`

	// **参数解释**: 命名空间名称 **取值范围**: 字符长度0-256位
	NamespaceName *string `json:"namespace_name,omitempty"`

	// **参数解释**: agent状态 **取值范围**: 字符长度0-65535位
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: agentID标识 **取值范围**: 字符长度0-65535位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 创建时间 **取值范围**: 最小值0，最大值9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**: 开通防护时间 **取值范围**: 最小值0，最大值9223372036854775807
	SubscribeTime *int64 `json:"subscribe_time,omitempty"`

	// **参数解释**: cpu申请值 **取值范围**: 字符长度0-65535位
	CpuRequest *string `json:"cpu_request,omitempty"`

	// **参数解释**: cpu限制值 **取值范围**: 字符长度0-65535位
	CpuLimit *string `json:"cpu_limit,omitempty"`

	// **参数解释**: 防护状态 **取值范围**: 字符长度0-65535位
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**: 策略组名称 **取值范围**: 字符长度0-65535位
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// **参数解释**: 标签 **取值范围**: 最小值0，最大值65535
	Labels *[]LabelInfo `json:"labels,omitempty"`

	// **参数解释**: 所属集群名称 **取值范围**: 字符长度0-65535位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**: 所属工作负载名称 **取值范围**: 字符长度0-65535位
	WorkloadName *string `json:"workload_name,omitempty"`

	// **参数解释**: 所属集群ID **取值范围**: 字符长度0-65535位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 容器列表 **取值范围**: 最小值0，最大值65535
	ContainerInfo *[]ProtectionContainerInfo `json:"container_info,omitempty"`

	// **参数解释**: 实例是否支持开关防护 **取值范围**: - true：是。 - false：否。
	IsSupportProtect *bool `json:"is_support_protect,omitempty"`

	// **参数解释**: 实例不支持开关防护的原因 **取值范围**: 字符长度0-65535位
	Reason *string `json:"reason,omitempty"`

	// **参数解释**: 实例漏洞风险数 **取值范围**: 最小值0，最大值9223372036854775807
	VulNum *int64 `json:"vul_num,omitempty"`

	// **参数解释**: 实例应用漏洞风险数 **取值范围**: 最小值0，最大值9223372036854775807
	AppVulNum *int64 `json:"app_vul_num,omitempty"`

	// **参数解释**: 实例系统漏洞风险数 **取值范围**: 最小值0，最大值9223372036854775807
	LinuxVulNum *int64 `json:"linux_vul_num,omitempty"`

	// **参数解释**: 实例基线检查数 **取值范围**: 最小值0，最大值9223372036854775807
	SecurityNum *int64 `json:"security_num,omitempty"`

	// **参数解释**: 实例入侵风险数 **取值范围**: 最小值0，最大值9223372036854775807
	EventNum *int64 `json:"event_num,omitempty"`
}

func (o ServerlessInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerlessInstanceInfo struct{}"
	}

	return strings.Join([]string{"ServerlessInstanceInfo", string(data)}, " ")
}
