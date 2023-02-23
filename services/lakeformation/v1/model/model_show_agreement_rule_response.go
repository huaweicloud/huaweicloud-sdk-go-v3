package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAgreementRuleResponse struct {

	// 系统协议列表
	AgreementRules *[]AgreementRule `json:"agreement_rules,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAgreementRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgreementRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowAgreementRuleResponse", string(data)}, " ")
}
