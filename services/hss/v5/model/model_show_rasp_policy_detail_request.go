package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRaspPolicyDetailRequest Request Object
type ShowRaspPolicyDetailRequest struct {

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 策略ID
	PolicyId string `json:"policy_id"`
}

func (o ShowRaspPolicyDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRaspPolicyDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowRaspPolicyDetailRequest", string(data)}, " ")
}
