package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFullRuleResponse Response Object
type UpdateFullRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateFullRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFullRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateFullRuleResponse", string(data)}, " ")
}
