package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowL7RuleResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	Rule           *L7Rule `json:"rule,omitempty" xml:"rule"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowL7RuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7RuleResponse struct{}"
	}

	return strings.Join([]string{"ShowL7RuleResponse", string(data)}, " ")
}
