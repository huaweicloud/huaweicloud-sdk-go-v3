package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2AgreementRuleRequest Request Object
type ListV2AgreementRuleRequest struct {
}

func (o ListV2AgreementRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2AgreementRuleRequest struct{}"
	}

	return strings.Join([]string{"ListV2AgreementRuleRequest", string(data)}, " ")
}
