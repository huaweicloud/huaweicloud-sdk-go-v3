package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePolicyRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
}

func (o DeletePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyRequest", string(data)}, " ")
}
