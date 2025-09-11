package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSqlExcuteNodesRequest Request Object
type ListSqlExcuteNodesRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListSqlExcuteNodesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *ListSqlExcuteNodeRequestBody `json:"body,omitempty"`
}

func (o ListSqlExcuteNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlExcuteNodesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlExcuteNodesRequest", string(data)}, " ")
}

type ListSqlExcuteNodesRequestXLanguage struct {
	value string
}

type ListSqlExcuteNodesRequestXLanguageEnum struct {
	ZH_CN ListSqlExcuteNodesRequestXLanguage
	EN_US ListSqlExcuteNodesRequestXLanguage
}

func GetListSqlExcuteNodesRequestXLanguageEnum() ListSqlExcuteNodesRequestXLanguageEnum {
	return ListSqlExcuteNodesRequestXLanguageEnum{
		ZH_CN: ListSqlExcuteNodesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSqlExcuteNodesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSqlExcuteNodesRequestXLanguage) Value() string {
	return c.value
}

func (c ListSqlExcuteNodesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlExcuteNodesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
