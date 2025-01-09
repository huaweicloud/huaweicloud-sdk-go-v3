package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableRuleRestrictionResponse Response Object
type EnableRuleRestrictionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableRuleRestrictionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableRuleRestrictionResponse struct{}"
	}

	return strings.Join([]string{"EnableRuleRestrictionResponse", string(data)}, " ")
}
