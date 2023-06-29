package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CustomConditions struct {

	// 字段类型。可选值为：url、user-agent、ip、params、cookie、referer、header、request_line、method、request
	Category *CustomConditionsCategory `json:"category,omitempty"`

	// 子字段：  - 字段类型为url、user-agent、ip、refer、request_line、method、request时，不需要传index参数    - 字段类型为params、header、cookie并且子字段为自定义时，index的值为自定义子字段
	Index *string `json:"index,omitempty"`

	// 条件列表匹配逻辑。   -  如果字段类型category是url、user-agent或者referer， 匹配逻辑可以为：contain、 not_contain、 equal、 not_equal、 prefix、 not_prefix、 suffix、 not_suffix、 contain_any、 not_contain_all、 equal_any、 not_equal_all、 equal_any、 not_equal_all、 prefix_any、 not_prefix_all、 suffix_any、 not_suffix_all、 len_greater、 len_less、len_equal或者len_not_equal   - 如果字段类型category是ip, 匹配逻辑可以为： equal、not_equal、equal_any或者not_equal_all   - 如果字段类型category是method, 匹配逻辑可以为： equal或者not_equal n - 如果字段类型category是request_line或者request, 匹配逻辑可以为： len_greater、len_less、len_equal或者len_not_equal   - 如果字段类型category是params、cookie或者header, 匹配逻辑可以为：contain、 not_contain、 equal、 not_equal、 prefix、 not_prefix、 suffix、 not_suffix、 contain_any、 not_contain_all、 equal_any、 not_equal_all、 equal_any、 not_equal_all、 prefix_any、 not_prefix_all、 suffix_any、 not_suffix_all、 len_greater、 len_less、len_equal、len_not_equal、、num_greater、num_less、num_equal、num_not_equal、exist或者not_exist
	LogicOperation *string `json:"logic_operation,omitempty"`

	// 条件列表逻辑匹配内容。当logic_operation参数不以any或者all结尾时，需要传该参数。
	Contents *[]string `json:"contents,omitempty"`

	// 引用表id。当logic_operation参数以any或者all结尾时，需要传该参数。此外，引用表类型要与category类型保持一致。
	ValueListId *string `json:"value_list_id,omitempty"`
}

func (o CustomConditions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomConditions struct{}"
	}

	return strings.Join([]string{"CustomConditions", string(data)}, " ")
}

type CustomConditionsCategory struct {
	value string
}

type CustomConditionsCategoryEnum struct {
	URL          CustomConditionsCategory
	USER_AGENT   CustomConditionsCategory
	REFERER      CustomConditionsCategory
	IP           CustomConditionsCategory
	METHOD       CustomConditionsCategory
	REQUEST_LINE CustomConditionsCategory
	REQUEST      CustomConditionsCategory
	PARAMS       CustomConditionsCategory
	COOKIE       CustomConditionsCategory
	HEADER       CustomConditionsCategory
}

func GetCustomConditionsCategoryEnum() CustomConditionsCategoryEnum {
	return CustomConditionsCategoryEnum{
		URL: CustomConditionsCategory{
			value: "url",
		},
		USER_AGENT: CustomConditionsCategory{
			value: "user-agent",
		},
		REFERER: CustomConditionsCategory{
			value: "referer",
		},
		IP: CustomConditionsCategory{
			value: "ip",
		},
		METHOD: CustomConditionsCategory{
			value: "method",
		},
		REQUEST_LINE: CustomConditionsCategory{
			value: "request_line",
		},
		REQUEST: CustomConditionsCategory{
			value: "request",
		},
		PARAMS: CustomConditionsCategory{
			value: "params",
		},
		COOKIE: CustomConditionsCategory{
			value: "cookie",
		},
		HEADER: CustomConditionsCategory{
			value: "header",
		},
	}
}

func (c CustomConditionsCategory) Value() string {
	return c.value
}

func (c CustomConditionsCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomConditionsCategory) UnmarshalJSON(b []byte) error {
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
