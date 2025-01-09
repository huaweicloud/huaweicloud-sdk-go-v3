package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRestrictedRuleResponse Response Object
type AddRestrictedRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddRestrictedRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRestrictedRuleResponse struct{}"
	}

	return strings.Join([]string{"AddRestrictedRuleResponse", string(data)}, " ")
}
