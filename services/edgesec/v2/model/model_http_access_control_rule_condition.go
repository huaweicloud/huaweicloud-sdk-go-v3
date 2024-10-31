package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpAccessControlRuleCondition 条件列表参数较为复杂，存在级联关系，建议同时使用控制台上的添加误报屏蔽规则，单击F12键查看路径后缀为/access-control-rule，方法为POST的请求参数，便于理解参数的填写
type HttpAccessControlRuleCondition struct {

	// 字段类型。可选值为：url、custom_asn、custom_geoip、robot、user-agent、ip、params、cookie、referer、header、method、request_line、request、response_code、response_length、response_time、response_header、response_body
	Category *string `json:"category,omitempty"`

	// 子字段：  - 字段类型为url、custom_asn、custom_geoip、robot、user-agent、referer、request_line、method、request、response_code、response_length、response_time、response_body时，不需要传index参数    - 字段类型为params、cookie、header、response_header并且子字段为自定义时，index的值为自定义子字段
	Index *string `json:"index,omitempty"`

	// 内容列表
	Contents *[]string `json:"contents,omitempty"`

	// 处理逻辑
	LogicOperation *string `json:"logic_operation,omitempty"`

	// 引用表id
	ValueListId *string `json:"value_list_id,omitempty"`

	// 若防护规则涉及阈值，即使用该字段
	Size *int64 `json:"size,omitempty"`

	// 1.所有子字段/2.任意子字段
	CheckAllIndexesLogic *int32 `json:"check_all_indexes_logic,omitempty"`
}

func (o HttpAccessControlRuleCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpAccessControlRuleCondition struct{}"
	}

	return strings.Join([]string{"HttpAccessControlRuleCondition", string(data)}, " ")
}
