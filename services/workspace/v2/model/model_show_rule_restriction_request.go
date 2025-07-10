package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRuleRestrictionRequest Request Object
type ShowRuleRestrictionRequest struct {
}

func (o ShowRuleRestrictionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRuleRestrictionRequest struct{}"
	}

	return strings.Join([]string{"ShowRuleRestrictionRequest", string(data)}, " ")
}
