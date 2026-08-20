package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListParameterGroupTemplatesRequest Request Object
type ListParameterGroupTemplatesRequest struct {

	// **参数解释**: 指定接口返回信息的语言类型。 **约束限制**: 不涉及。 **取值范围**: - zh-cn：中文 - en-us：英文  **默认取值**: en-us
	XLanguage *ListParameterGroupTemplatesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。例如：该参数指定为0，limit指定为10，则只展示第1~10条数据。 **约束限制**: 不涉及。 **取值范围**: [0, 2147483647] **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 查询记录数。例如该参数设定为10，则查询结果最多只显示10条记录。 **约束限制**: 不涉及。 **取值范围**: [1, 100] **默认取值**: 默认为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListParameterGroupTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListParameterGroupTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListParameterGroupTemplatesRequest", string(data)}, " ")
}

type ListParameterGroupTemplatesRequestXLanguage struct {
	value string
}

type ListParameterGroupTemplatesRequestXLanguageEnum struct {
	ZH_CN ListParameterGroupTemplatesRequestXLanguage
	EN_US ListParameterGroupTemplatesRequestXLanguage
}

func GetListParameterGroupTemplatesRequestXLanguageEnum() ListParameterGroupTemplatesRequestXLanguageEnum {
	return ListParameterGroupTemplatesRequestXLanguageEnum{
		ZH_CN: ListParameterGroupTemplatesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListParameterGroupTemplatesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListParameterGroupTemplatesRequestXLanguage) Value() string {
	return c.value
}

func (c ListParameterGroupTemplatesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListParameterGroupTemplatesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
