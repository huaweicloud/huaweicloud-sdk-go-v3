package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpAccessControlRuleRequest Request Object
type CreateHttpAccessControlRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *CreateHttpAccessControlRuleRequestBody `json:"body,omitempty"`
}

func (o CreateHttpAccessControlRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpAccessControlRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateHttpAccessControlRuleRequest", string(data)}, " ")
}
