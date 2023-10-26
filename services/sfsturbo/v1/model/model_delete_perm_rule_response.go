package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePermRuleResponse Response Object
type DeletePermRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePermRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePermRuleResponse struct{}"
	}

	return strings.Join([]string{"DeletePermRuleResponse", string(data)}, " ")
}
