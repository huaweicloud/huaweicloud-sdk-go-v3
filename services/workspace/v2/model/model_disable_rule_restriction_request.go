package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableRuleRestrictionRequest Request Object
type DisableRuleRestrictionRequest struct {
}

func (o DisableRuleRestrictionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableRuleRestrictionRequest struct{}"
	}

	return strings.Join([]string{"DisableRuleRestrictionRequest", string(data)}, " ")
}
