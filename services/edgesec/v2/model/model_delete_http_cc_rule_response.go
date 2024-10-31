package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpCcRuleResponse Response Object
type DeleteHttpCcRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHttpCcRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpCcRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteHttpCcRuleResponse", string(data)}, " ")
}
