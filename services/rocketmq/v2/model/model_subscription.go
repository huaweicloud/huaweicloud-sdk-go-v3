package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Subscription struct {

	// 订阅的Topic名称。
	Topic *string `json:"topic,omitempty"`

	// 订阅类型，取值如下：TAG和SQL92。
	Type *string `json:"type,omitempty"`

	// 订阅tag字符。
	Expression *string `json:"expression,omitempty"`
}

func (o Subscription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subscription struct{}"
	}

	return strings.Join([]string{"Subscription", string(data)}, " ")
}
