package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateL7RuleResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	Rule           *L7Rule `json:"rule,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateL7RuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7RuleResponse struct{}"
	}

	return strings.Join([]string{"CreateL7RuleResponse", string(data)}, " ")
}
