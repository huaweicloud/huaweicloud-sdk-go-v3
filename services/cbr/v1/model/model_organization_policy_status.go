package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationPolicyStatus 组织策略策略部署状态
type OrganizationPolicyStatus struct {

	// 策略ID
	PolicyId string `json:"policy_id"`

	// 账号ID
	DomainId string `json:"domain_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 状态
	Status string `json:"status"`
}

func (o OrganizationPolicyStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationPolicyStatus struct{}"
	}

	return strings.Join([]string{"OrganizationPolicyStatus", string(data)}, " ")
}
