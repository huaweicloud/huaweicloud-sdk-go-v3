package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroupPolicyResponseInfo 安全组策略
type SecurityGroupPolicyResponseInfo struct {

	// 策略Id
	PolicyId string `json:"policy_id"`

	// 策略名称
	PolicyName string `json:"policy_name"`

	// 选择器
	PodSelector string `json:"pod_selector"`

	// 安全组列表
	SecurityGroups []SecurityGroup `json:"security_groups"`

	// 工作负载ID
	WorkloadId string `json:"workload_id"`

	// 工作负载名称
	WorkloadName string `json:"workload_name"`

	// 工作负载类型
	WorkloadType string `json:"workload_type"`

	// 命名空间id
	NamespaceId *string `json:"namespace_id,omitempty"`

	// 命名空间
	Namespace string `json:"namespace"`

	// 创建时间
	CreateTime string `json:"create_time"`

	// 策略yaml格式内容
	PolicyContent string `json:"policy_content"`
}

func (o SecurityGroupPolicyResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupPolicyResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityGroupPolicyResponseInfo", string(data)}, " ")
}
