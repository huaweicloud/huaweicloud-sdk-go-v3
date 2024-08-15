package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateL7RuleResponse Response Object
type UpdateL7RuleResponse struct {

	// 参数解释：请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	Rule           *L7Rule `json:"rule,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateL7RuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7RuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateL7RuleResponse", string(data)}, " ")
}
