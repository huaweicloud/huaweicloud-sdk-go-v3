package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateGeoipRuleRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// geoipRuleId

	RuleId string `json:"rule_id"`

	Body *UpdateGeoipRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateGeoipRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGeoipRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateGeoipRuleRequest", string(data)}, " ")
}
