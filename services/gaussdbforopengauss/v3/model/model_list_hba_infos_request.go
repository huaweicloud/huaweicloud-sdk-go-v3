package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHbaInfosRequest Request Object
type ListHbaInfosRequest struct {

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListHbaInfosRequestXLanguage `json:"X-Language,omitempty"`

	// **参数描述**: 分页符。从第一条数据偏移offset页数据后开始查询。例如，该参数指定为1，limit指定为10，则只展示第11-20条数据。 **约束限制**: 不涉及。 **取值范围**: [0, 2^31-1] **默认值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	// **参数描述**: 每页显示的条目数量。 **约束限制**: 不涉及。 **取值范围**: [1, 100] **默认值**: 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListHbaInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHbaInfosRequest struct{}"
	}

	return strings.Join([]string{"ListHbaInfosRequest", string(data)}, " ")
}

type ListHbaInfosRequestXLanguage struct {
	value string
}

type ListHbaInfosRequestXLanguageEnum struct {
	ZH_CN ListHbaInfosRequestXLanguage
	EN_US ListHbaInfosRequestXLanguage
}

func GetListHbaInfosRequestXLanguageEnum() ListHbaInfosRequestXLanguageEnum {
	return ListHbaInfosRequestXLanguageEnum{
		ZH_CN: ListHbaInfosRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListHbaInfosRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListHbaInfosRequestXLanguage) Value() string {
	return c.value
}

func (c ListHbaInfosRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHbaInfosRequestXLanguage) UnmarshalJSON(b []byte) error {
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
