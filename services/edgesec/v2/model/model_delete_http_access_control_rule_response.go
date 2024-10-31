package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpAccessControlRuleResponse Response Object
type DeleteHttpAccessControlRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHttpAccessControlRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpAccessControlRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteHttpAccessControlRuleResponse", string(data)}, " ")
}
