package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgreementRuleRequest Request Object
type ShowAgreementRuleRequest struct {
}

func (o ShowAgreementRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgreementRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowAgreementRuleRequest", string(data)}, " ")
}
