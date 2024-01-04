package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSubscriptionsFilterPolicesRequestBody filter_polices字段，必填，一个筛选策略最多可具有 5 个属性策略。  name-属性名称：满足正则 ^(?!smn_)(?!.*?_$)(?!^_)(?!.*?__)[a-z0-9_]{1,32}$ （英文字母([a-z])、数字([0-9])、下划线(_)，下划线不得开始、结尾或连续出现），长度限制{1,32}，key不能是smn_开头 字符串匹配策略： string_equals：精准匹配字符串；类型：字符串数组，数组长度[1,10]，字符串长度限制[1,32]，不能重复，不能空值“ ”
type BatchSubscriptionsFilterPolicesRequestBody struct {

	// 批量处理订阅者策略列表。
	Polices []BatchSubscriptionsFilterPolicesRequestBodyPolices `json:"polices"`
}

func (o BatchSubscriptionsFilterPolicesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSubscriptionsFilterPolicesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSubscriptionsFilterPolicesRequestBody", string(data)}, " ")
}
