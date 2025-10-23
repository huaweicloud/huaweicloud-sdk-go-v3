package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSqlPlanActionRequest Request Object
type ListSqlPlanActionRequest struct {

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 节点ID，仅支持非日志类型的CN或主DN、备DN节点。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListSqlPlanActionRequestXLanguage `json:"X-Language,omitempty"`

	Body *QuerySqlPlanStateRequest `json:"body,omitempty"`
}

func (o ListSqlPlanActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlPlanActionRequest struct{}"
	}

	return strings.Join([]string{"ListSqlPlanActionRequest", string(data)}, " ")
}

type ListSqlPlanActionRequestXLanguage struct {
	value string
}

type ListSqlPlanActionRequestXLanguageEnum struct {
	ZH_CN ListSqlPlanActionRequestXLanguage
	EN_US ListSqlPlanActionRequestXLanguage
}

func GetListSqlPlanActionRequestXLanguageEnum() ListSqlPlanActionRequestXLanguageEnum {
	return ListSqlPlanActionRequestXLanguageEnum{
		ZH_CN: ListSqlPlanActionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSqlPlanActionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSqlPlanActionRequestXLanguage) Value() string {
	return c.value
}

func (c ListSqlPlanActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlPlanActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
