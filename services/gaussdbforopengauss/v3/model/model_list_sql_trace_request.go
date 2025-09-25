package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSqlTraceRequest Request Object
type ListSqlTraceRequest struct {

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 归一化SQL ID，对应内核字段：unique_sql_id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: 唯一SQL ID，对应内核字段：debug_query_id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlExecId *string `json:"sql_exec_id,omitempty"`

	// **参数解释**: 事务ID，对应内核字段：transaction_id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TransactionId *string `json:"transaction_id,omitempty"`

	// **参数解释**: 链路ID，对应内核字段：trace_id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListSqlTraceRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListSqlTraceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlTraceRequest struct{}"
	}

	return strings.Join([]string{"ListSqlTraceRequest", string(data)}, " ")
}

type ListSqlTraceRequestXLanguage struct {
	value string
}

type ListSqlTraceRequestXLanguageEnum struct {
	ZH_CN ListSqlTraceRequestXLanguage
	EN_US ListSqlTraceRequestXLanguage
}

func GetListSqlTraceRequestXLanguageEnum() ListSqlTraceRequestXLanguageEnum {
	return ListSqlTraceRequestXLanguageEnum{
		ZH_CN: ListSqlTraceRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSqlTraceRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSqlTraceRequestXLanguage) Value() string {
	return c.value
}

func (c ListSqlTraceRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlTraceRequestXLanguage) UnmarshalJSON(b []byte) error {
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
