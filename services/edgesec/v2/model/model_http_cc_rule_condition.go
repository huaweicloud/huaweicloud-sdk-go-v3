package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HttpCcRuleCondition 防护规则条件
type HttpCcRuleCondition struct {

	// 防护规则字段
	Category HttpCcRuleConditionCategory `json:"category"`

	// 子字段，当字段类型（category）选择“params”、“cookie”、“header”时，请根据实际需求配置子字段且该参数必填。
	Index *string `json:"index,omitempty"`

	// 条件列表逻辑匹配内容。当logic_operation参数不以any或者all结尾时，需要传该参数。
	Contents *[]string `json:"contents,omitempty"`

	// 条件列表匹配逻辑。   -  如果字段类型category是url， 匹配逻辑可以为：contain、 not_contain、 equal、 not_equal、 prefix、 not_prefix、 suffix、 not_suffix、 contain_any、 not_contain_all、 equal_any、 not_equal_all、 equal_any、 not_equal_all、 prefix_any、 not_prefix_all、 suffix_any、 not_suffix_all、 len_greater、 len_less、len_equal或者len_not_equal   - 如果字段类型category是ip或者ipv6，匹配逻辑可以为： equal、not_equal、equal_any或者not_equal_all   - 如果字段类型category是params、cookie或者header, 匹配逻辑可以为：contain、 not_contain、 equal、 not_equal、 prefix、 not_prefix、 suffix、 not_suffix、 contain_any、 not_contain_all、 equal_any、not_equal_all、 equal_any、 not_equal_all、 prefix_any、 not_prefix_all、 suffix_any、 not_suffix_all、 len_greater、 len_less、len_equal、len_not_equal、、num_greater、num_less、num_equal、num_not_equal、exist或者not_exist
	LogicOperation string `json:"logic_operation"`

	// 引用表id。当logic_operation参数以any或者all结尾时，需要传该参数。此外，引用表类型要与category类型保持一致。
	ValueListId *string `json:"value_list_id,omitempty"`

	// 若防护规则涉及阈值，即使用该字段
	Size *int64 `json:"size,omitempty"`

	// - 1：所有子字段 - 2：任意子字段
	CheckAllIndexesLogic *int32 `json:"check_all_indexes_logic,omitempty"`
}

func (o HttpCcRuleCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpCcRuleCondition struct{}"
	}

	return strings.Join([]string{"HttpCcRuleCondition", string(data)}, " ")
}

type HttpCcRuleConditionCategory struct {
	value string
}

type HttpCcRuleConditionCategoryEnum struct {
	URL           HttpCcRuleConditionCategory
	IP            HttpCcRuleConditionCategory
	PARAMS        HttpCcRuleConditionCategory
	COOKIE        HttpCcRuleConditionCategory
	HEADER        HttpCcRuleConditionCategory
	CUSTOM_GEOIP  HttpCcRuleConditionCategory
	RESPONSE_CODE HttpCcRuleConditionCategory
}

func GetHttpCcRuleConditionCategoryEnum() HttpCcRuleConditionCategoryEnum {
	return HttpCcRuleConditionCategoryEnum{
		URL: HttpCcRuleConditionCategory{
			value: "url",
		},
		IP: HttpCcRuleConditionCategory{
			value: "ip",
		},
		PARAMS: HttpCcRuleConditionCategory{
			value: "params",
		},
		COOKIE: HttpCcRuleConditionCategory{
			value: "cookie",
		},
		HEADER: HttpCcRuleConditionCategory{
			value: "header",
		},
		CUSTOM_GEOIP: HttpCcRuleConditionCategory{
			value: "custom_geoip",
		},
		RESPONSE_CODE: HttpCcRuleConditionCategory{
			value: "response_code",
		},
	}
}

func (c HttpCcRuleConditionCategory) Value() string {
	return c.value
}

func (c HttpCcRuleConditionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HttpCcRuleConditionCategory) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
