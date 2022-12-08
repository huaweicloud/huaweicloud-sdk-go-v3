package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ChangeRuleResponse struct {

	// 返回消息
	Msg *string `json:"msg,omitempty"`

	// 返回状态，如'200','400'
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeRuleResponse struct{}"
	}

	return strings.Join([]string{"ChangeRuleResponse", string(data)}, " ")
}
