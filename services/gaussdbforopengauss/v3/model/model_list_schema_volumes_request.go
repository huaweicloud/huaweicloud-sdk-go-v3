package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSchemaVolumesRequest Request Object
type ListSchemaVolumesRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListSchemaVolumesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 数据库名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	DatabaseName string `json:"database_name"`

	// **参数解释**: schema的名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。 **约束限制**: 必须为数字，不能为负数。 **取值范围**: 0 - 10000 **默认取值**: 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 查询记录数。 **约束限制**: 不能为负数 **取值范围**: 最小值为1，最大值为200。 **默认取值**: 默认为200。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSchemaVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemaVolumesRequest struct{}"
	}

	return strings.Join([]string{"ListSchemaVolumesRequest", string(data)}, " ")
}

type ListSchemaVolumesRequestXLanguage struct {
	value string
}

type ListSchemaVolumesRequestXLanguageEnum struct {
	ZH_CN ListSchemaVolumesRequestXLanguage
	EN_US ListSchemaVolumesRequestXLanguage
}

func GetListSchemaVolumesRequestXLanguageEnum() ListSchemaVolumesRequestXLanguageEnum {
	return ListSchemaVolumesRequestXLanguageEnum{
		ZH_CN: ListSchemaVolumesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSchemaVolumesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSchemaVolumesRequestXLanguage) Value() string {
	return c.value
}

func (c ListSchemaVolumesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSchemaVolumesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
