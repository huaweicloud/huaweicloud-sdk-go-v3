package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscriptionsFilterPolicy 过滤策略。
type SubscriptionsFilterPolicy struct {

	// 过滤策略名称。 包含小写英文字母([a-z])、数字([0-9])、下划线(_)，下划线不得开始、结尾或连续出现），长度限制{1,32}，不能是smn_开头。
	Name string `json:"name"`

	// 字符串精确匹配数组。数组长度[1, 10]，数组内容不能重复，值不能为null或者空字符串“ ”，长度限制[1,32]，中英文、数字、下划线
	StringEquals []string `json:"string_equals"`
}

func (o SubscriptionsFilterPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionsFilterPolicy struct{}"
	}

	return strings.Join([]string{"SubscriptionsFilterPolicy", string(data)}, " ")
}
