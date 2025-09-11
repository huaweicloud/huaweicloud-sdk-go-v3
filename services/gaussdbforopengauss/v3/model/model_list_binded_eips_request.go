package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBindedEipsRequest Request Object
type ListBindedEipsRequest struct {

	// 语言
	XLanguage *ListBindedEipsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为50，不能为负数，最小值为1，最大值为50。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListBindedEipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBindedEipsRequest struct{}"
	}

	return strings.Join([]string{"ListBindedEipsRequest", string(data)}, " ")
}

type ListBindedEipsRequestXLanguage struct {
	value string
}

type ListBindedEipsRequestXLanguageEnum struct {
	ZH_CN ListBindedEipsRequestXLanguage
	EN_US ListBindedEipsRequestXLanguage
}

func GetListBindedEipsRequestXLanguageEnum() ListBindedEipsRequestXLanguageEnum {
	return ListBindedEipsRequestXLanguageEnum{
		ZH_CN: ListBindedEipsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListBindedEipsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListBindedEipsRequestXLanguage) Value() string {
	return c.value
}

func (c ListBindedEipsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBindedEipsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
