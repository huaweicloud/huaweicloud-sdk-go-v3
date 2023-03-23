package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomRuleConditions struct {

	// 字段类型。可选值为：url、user-agent、ip、params、cookie、referer、header、request_line、method、reqeust
	Category *string `json:"category,omitempty"`

	// 子字段：  - 字段类型为url、user-agent、ip、refer、request_line、method、reqeust时，不需要传index参数    - 字段类型为params、header、cookie并且子字段为自定义时，index的值为自定义子字段
	Index *string `json:"index,omitempty"`

	// 条件匹配逻辑。
	LogicOperation *string `json:"logic_operation,omitempty"`

	// 条件匹配的内容
	Contents *[]string `json:"contents,omitempty"`

	// 引用表id。
	ValueListId *string `json:"value_list_id,omitempty"`
}

func (o CustomRuleConditions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomRuleConditions struct{}"
	}

	return strings.Join([]string{"CustomRuleConditions", string(data)}, " ")
}
