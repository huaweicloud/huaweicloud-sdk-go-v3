package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyRequest Request Object
type ShowPolicyRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 策略Id
	PolicyId string `json:"policy_id"`
}

func (o ShowPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyRequest", string(data)}, " ")
}
