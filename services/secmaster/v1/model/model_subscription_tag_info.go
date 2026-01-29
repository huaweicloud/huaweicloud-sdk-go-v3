package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionTagInfo struct {

	// 键。 最大长度36个字符。 字符集：A-Z，a-z ， 0-9，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）
	Key *string `json:"key,omitempty"`

	// 值。 最大长度43个字符，可以为空字符串。 字符集：A-Z，a-z ， 0-9，‘.’，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）
	Value *string `json:"value,omitempty"`

	// 创建时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o SubscriptionTagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionTagInfo struct{}"
	}

	return strings.Join([]string{"SubscriptionTagInfo", string(data)}, " ")
}
