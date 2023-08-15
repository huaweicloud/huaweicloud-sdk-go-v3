package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddActionRuleResponse Response Object
type AddActionRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddActionRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddActionRuleResponse struct{}"
	}

	return strings.Join([]string{"AddActionRuleResponse", string(data)}, " ")
}
