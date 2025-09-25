package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFullSqlSwitchesRequest Request Object
type ListFullSqlSwitchesRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListFullSqlSwitchesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 查询记录数。 **约束限制**: 不涉及。 **取值范围**: 1~1000。 **默认取值**: 默认为100。
	Limit int32 `json:"limit"`

	// **参数解释**: 索引位置，偏移量。 **约束限制**: 不涉及。 **取值范围**: 0 ~ 2,147,483,647。 **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset int32 `json:"offset"`
}

func (o ListFullSqlSwitchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFullSqlSwitchesRequest struct{}"
	}

	return strings.Join([]string{"ListFullSqlSwitchesRequest", string(data)}, " ")
}

type ListFullSqlSwitchesRequestXLanguage struct {
	value string
}

type ListFullSqlSwitchesRequestXLanguageEnum struct {
	ZH_CN ListFullSqlSwitchesRequestXLanguage
	EN_US ListFullSqlSwitchesRequestXLanguage
}

func GetListFullSqlSwitchesRequestXLanguageEnum() ListFullSqlSwitchesRequestXLanguageEnum {
	return ListFullSqlSwitchesRequestXLanguageEnum{
		ZH_CN: ListFullSqlSwitchesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListFullSqlSwitchesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListFullSqlSwitchesRequestXLanguage) Value() string {
	return c.value
}

func (c ListFullSqlSwitchesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFullSqlSwitchesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
