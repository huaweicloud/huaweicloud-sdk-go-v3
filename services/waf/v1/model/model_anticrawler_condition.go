package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AnticrawlerCondition struct {

	// 字段类型
	Category *AnticrawlerConditionCategory `json:"category,omitempty"`

	// 条件列表匹配逻辑, 包括：contain、not_contain、equal、not_equal、prefix、not_prefix、suffix、not_suffix、contain_any、not_contain_all、equal_any、not_equal_all、prefix_any、not_prefix_all、suffix_any、not_suffix_all
	LogicOperation *string `json:"logic_operation,omitempty"`

	// 条件列表逻辑匹配内容。当logic_operation参数不以any或者all结尾时，需要传该参数。
	Contents *[]string `json:"contents,omitempty"`

	// 引用表id。当logic_operation参数以any或者all结尾时，需要传该参数。此外，引用表类型要与category类型保持一致。
	ValueListId *string `json:"value_list_id,omitempty"`
}

func (o AnticrawlerCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnticrawlerCondition struct{}"
	}

	return strings.Join([]string{"AnticrawlerCondition", string(data)}, " ")
}

type AnticrawlerConditionCategory struct {
	value string
}

type AnticrawlerConditionCategoryEnum struct {
	URL        AnticrawlerConditionCategory
	USER_AGENT AnticrawlerConditionCategory
}

func GetAnticrawlerConditionCategoryEnum() AnticrawlerConditionCategoryEnum {
	return AnticrawlerConditionCategoryEnum{
		URL: AnticrawlerConditionCategory{
			value: "url",
		},
		USER_AGENT: AnticrawlerConditionCategory{
			value: "user-agent",
		},
	}
}

func (c AnticrawlerConditionCategory) Value() string {
	return c.value
}

func (c AnticrawlerConditionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnticrawlerConditionCategory) UnmarshalJSON(b []byte) error {
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
