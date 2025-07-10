package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRuleRestrictionRequest Request Object
type SetRuleRestrictionRequest struct {
	Body *SetRuleRestrictionReq `json:"body,omitempty"`
}

func (o SetRuleRestrictionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRuleRestrictionRequest struct{}"
	}

	return strings.Join([]string{"SetRuleRestrictionRequest", string(data)}, " ")
}
