package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRuleRestrictionResponse Response Object
type SetRuleRestrictionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetRuleRestrictionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRuleRestrictionResponse struct{}"
	}

	return strings.Join([]string{"SetRuleRestrictionResponse", string(data)}, " ")
}
