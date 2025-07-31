package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSecurityGroupPolicyRequestBody struct {

	// 策略名称
	PolicyName string `json:"policy_name"`

	// 工作负载ID
	WorkloadId string `json:"workload_id"`

	// 工作负载名称
	WorkloadName string `json:"workload_name"`

	// 工作负载类型
	WorkloadType string `json:"workload_type"`

	// 安全组列表
	SecurityGroups []SecurityGroup `json:"security_groups"`
}

func (o CreateSecurityGroupPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupPolicyRequestBody", string(data)}, " ")
}
