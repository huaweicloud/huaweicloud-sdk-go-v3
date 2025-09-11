package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMetricDatasRequest Request Object
type ListMetricDatasRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ListMetricDatasRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。例如：该参数指定为1，limit指定为10，则只展示第2~11条数据。 **约束限制**: 不涉及。 **取值范围**: [0, 2^31-1] **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 查询记录数。例如该参数设定为10，则查询结果最多只显示10条记录。 **约束限制**: 不涉及。 **取值范围**: [1, 50] **默认取值**: 默认为50。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 实例名称。 用于表示实例的名称，同一租户下，同类型的实例名可重名。 **约束限制**: 不涉及。 **取值范围**: 4~64个字符之间，必须以字母开头，区分大小写，可以包含字母、数字、中划线或者下划线，不能包含其他的特殊字符。 **默认取值**: 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**: 实例名称。 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ListMetricDatasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricDatasRequest struct{}"
	}

	return strings.Join([]string{"ListMetricDatasRequest", string(data)}, " ")
}

type ListMetricDatasRequestXLanguage struct {
	value string
}

type ListMetricDatasRequestXLanguageEnum struct {
	ZH_CN ListMetricDatasRequestXLanguage
	EN_US ListMetricDatasRequestXLanguage
}

func GetListMetricDatasRequestXLanguageEnum() ListMetricDatasRequestXLanguageEnum {
	return ListMetricDatasRequestXLanguageEnum{
		ZH_CN: ListMetricDatasRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListMetricDatasRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListMetricDatasRequestXLanguage) Value() string {
	return c.value
}

func (c ListMetricDatasRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMetricDatasRequestXLanguage) UnmarshalJSON(b []byte) error {
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
