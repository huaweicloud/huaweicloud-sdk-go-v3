package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateL7ruleResponse struct {
	Rule           *L7ruleResp `json:"rule,omitempty" xml:"rule"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateL7ruleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7ruleResponse struct{}"
	}

	return strings.Join([]string{"CreateL7ruleResponse", string(data)}, " ")
}
