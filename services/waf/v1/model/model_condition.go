package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Condition struct {

	// 字段类型，可选值有ip、url、params、cookie、header
	Category *string `json:"category,omitempty"`

	// 内容,数组长度限制为1，内容格式根据字段类型变化，例如，字段类型为ip时，contents内容格式需为ip地址或ip地址段；字段类型为url时，contents内容格式需为标准url格式；字段类型为params、cookie、header时，内容的格式不做限制
	Contents *[]string `json:"contents,omitempty"`

	// 匹配逻辑，匹配逻辑根据字段类型变化，字段类型为ip时，匹配逻辑支持（equal：等于，not_equal：不等于），字段类型为url、header、params、cookie时，匹配逻辑支持（equal：等于，not_equal：不等于，contain：包含，not_contain：不包含，prefix：前缀为，not_prefix：前缀不为，suffix：后缀为，not_suffix：后缀不为，regular_match：正则匹配，regular_not_match：正则不匹配）
	LogicOperation *string `json:"logic_operation,omitempty"`

	// 字段类型为url或ip时不存在check_all_indexes_logic字段，其它情况下（1：检查所有子字段，2：检查任意子字段，null：使用自定义子字段）
	CheckAllIndexesLogic *int32 `json:"check_all_indexes_logic,omitempty"`

	// 字段类型为ip且子字段为客户端ip时，不存在index参数；子字段类型为X-Forwarded-For时，值为x-forwarded-for，字段类型为params、header、cookie并且子字段为自定义时，index的值为自定义子字段
	Index *string `json:"index,omitempty"`
}

func (o Condition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Condition struct{}"
	}

	return strings.Join([]string{"Condition", string(data)}, " ")
}
