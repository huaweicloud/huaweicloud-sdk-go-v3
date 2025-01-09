package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRestrictedRuleResponse Response Object
type DeleteRestrictedRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRestrictedRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRestrictedRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteRestrictedRuleResponse", string(data)}, " ")
}
