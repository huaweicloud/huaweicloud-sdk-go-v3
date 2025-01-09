package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppRuleResponse Response Object
type DeleteAppRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppRuleResponse", string(data)}, " ")
}
