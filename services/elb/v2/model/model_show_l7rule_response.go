package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowL7ruleResponse struct {
	Rule           *L7ruleResp `json:"rule,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowL7ruleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7ruleResponse struct{}"
	}

	return strings.Join([]string{"ShowL7ruleResponse", string(data)}, " ")
}
