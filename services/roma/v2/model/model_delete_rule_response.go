package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRuleResponse Response Object
type DeleteRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteRuleResponse", string(data)}, " ")
}
