package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRuleGroupResponse Response Object
type DeleteRuleGroupResponse struct {

	// 返回消息
	Msg *string `json:"msg,omitempty"`

	// 返回状态，如'200','400'
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRuleGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteRuleGroupResponse", string(data)}, " ")
}
