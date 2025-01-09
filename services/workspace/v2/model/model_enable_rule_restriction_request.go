package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableRuleRestrictionRequest Request Object
type EnableRuleRestrictionRequest struct {
}

func (o EnableRuleRestrictionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableRuleRestrictionRequest struct{}"
	}

	return strings.Join([]string{"EnableRuleRestrictionRequest", string(data)}, " ")
}
