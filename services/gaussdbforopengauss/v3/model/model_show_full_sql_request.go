package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFullSqlRequest Request Object
type ShowFullSqlRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ShowFullSqlRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 唯一SQL ID，对应内核字段：debug_query_id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlExecId string `json:"sql_exec_id"`

	// **参数解释**: 采集到的全量SQL记录的唯一键ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释**: 归一化SQL ID，对应内核字段：unique_sql_id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`
}

func (o ShowFullSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFullSqlRequest struct{}"
	}

	return strings.Join([]string{"ShowFullSqlRequest", string(data)}, " ")
}

type ShowFullSqlRequestXLanguage struct {
	value string
}

type ShowFullSqlRequestXLanguageEnum struct {
	ZH_CN ShowFullSqlRequestXLanguage
	EN_US ShowFullSqlRequestXLanguage
}

func GetShowFullSqlRequestXLanguageEnum() ShowFullSqlRequestXLanguageEnum {
	return ShowFullSqlRequestXLanguageEnum{
		ZH_CN: ShowFullSqlRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowFullSqlRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowFullSqlRequestXLanguage) Value() string {
	return c.value
}

func (c ShowFullSqlRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFullSqlRequestXLanguage) UnmarshalJSON(b []byte) error {
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
