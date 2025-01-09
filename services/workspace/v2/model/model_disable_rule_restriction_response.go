package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableRuleRestrictionResponse Response Object
type DisableRuleRestrictionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableRuleRestrictionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableRuleRestrictionResponse struct{}"
	}

	return strings.Join([]string{"DisableRuleRestrictionResponse", string(data)}, " ")
}
