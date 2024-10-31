package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpIgnoreRuleResponse Response Object
type DeleteHttpIgnoreRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHttpIgnoreRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpIgnoreRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteHttpIgnoreRuleResponse", string(data)}, " ")
}
