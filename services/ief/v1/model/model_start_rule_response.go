package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartRuleResponse struct{}"
	}

	return strings.Join([]string{"StartRuleResponse", string(data)}, " ")
}
