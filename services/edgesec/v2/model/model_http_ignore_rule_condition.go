package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpIgnoreRuleCondition 条件列表参数较为复杂，存在级联关系，建议同时使用控制台上的添加误报屏蔽规则，单击F12键查看路径后缀为/ignore-rule，方法为POST的请求参数，便于理解参数的填写
type HttpIgnoreRuleCondition struct {

	// 字段类型，可选值有ip、url、params、cookie、header
	Category string `json:"category"`

	// 字段类型为ip且子字段为客户端ip时，不需要传index参数；子字段类型为X-Forwarded-For时，值为x-forwarded-for；字段类型为params、header、cookie并且子字段为自定义时，index的值为自定义子字段
	Index *string `json:"index,omitempty"`

	// 内容列表
	Contents []string `json:"contents"`

	// 匹配逻辑，匹配逻辑根据字段类型变化，字段类型为ip时，匹配逻辑支持（equal：等于，not_equal：不等于），字段类型为url、header、params、cookie时，匹配逻辑支持（equal：等于，not_equal：不等于，contain：包含，not_contain：不包含，prefix：前缀为，not_prefix：前缀不为，suffix：后缀为，not_suffix：后缀不为，regular_match：正则匹配，regular_not_match：正则不匹配）
	LogicOperation string `json:"logic_operation"`

	// 引用表id
	ValueListId *string `json:"value_list_id,omitempty"`

	// 若防护规则涉及阈值，即使用该字段
	Size *int64 `json:"size,omitempty"`

	// 1.所有子字段/2.任意子字段
	CheckAllIndexesLogic *int32 `json:"check_all_indexes_logic,omitempty"`
}

func (o HttpIgnoreRuleCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpIgnoreRuleCondition struct{}"
	}

	return strings.Join([]string{"HttpIgnoreRuleCondition", string(data)}, " ")
}
