package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSubscriptionUserRequestBody struct {

	// 订阅用户名称。
	Name *string `json:"name,omitempty"`

	// 订阅用户分组。每个订阅分组只能包含小写英文字母([a-z])、数字([0-9])、下划线(_)，下划线不能出现在开始或结尾，下划线不能连续出现，长度为1到32个字符。
	Group *[]string `json:"group,omitempty"`
}

func (o UpdateSubscriptionUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionUserRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionUserRequestBody", string(data)}, " ")
}
