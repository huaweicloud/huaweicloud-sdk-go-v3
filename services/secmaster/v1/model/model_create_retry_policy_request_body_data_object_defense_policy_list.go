package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRetryPolicyRequestBodyDataObjectDefensePolicyList 与操作连接对应的策略
type CreateRetryPolicyRequestBodyDataObjectDefensePolicyList struct {

	// 操作连接ID
	DefenseConnectionId string `json:"defense_connection_id"`

	// 操作连接名称
	DefenseConnectionName string `json:"defense_connection_name"`

	// 防线策略归属区域ID
	DefenseConnectionRegionId string `json:"defense_connection_region_id"`

	// 防线策略归属区域名称
	DefenseConnectionRegionName string `json:"defense_connection_region_name"`

	// 防线服务
	DefenseType string `json:"defense_type"`

	// 企业项目ID
	TargetEnterpriseId string `json:"target_enterprise_id"`

	// 企业项目名称
	TargetEnterpriseName string `json:"target_enterprise_name"`

	// 防线策略归属项目ID
	TargetProjectId string `json:"target_project_id"`

	// 防线策略归属项目名称
	TargetProjectName string `json:"target_project_name"`
}

func (o CreateRetryPolicyRequestBodyDataObjectDefensePolicyList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetryPolicyRequestBodyDataObjectDefensePolicyList struct{}"
	}

	return strings.Join([]string{"CreateRetryPolicyRequestBodyDataObjectDefensePolicyList", string(data)}, " ")
}
