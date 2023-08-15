package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRuleResponse Response Object
type DeleteRuleResponse struct {

	// 返回消息
	Msg *string `json:"msg,omitempty"`

	// 返回状态，如'200','400'
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteRuleResponse", string(data)}, " ")
}
