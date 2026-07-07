package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEventsRequest Request Object
type ListEventsRequest struct {

	// **参数解释**: 请求语言类型。 **约束限制**: 不涉及。 **取值范围**: - en-us - zh-cn  **默认取值**: en-us。
	XLanguage *ListEventsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 事件ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 事件状态。 **约束限制**: 不涉及。 **取值范围**: - WAITING：等待中 - INQUIRING：待授权 - SCHEDULED：待执行 - EXECUTING：执行中 - COMPLETED：已完成 - FAILED：失败 - CANCELED：已取消 **默认取值**: 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 事件类型。 **约束限制**: 不涉及。 **取值范围**: - RESTAT_NODE：重启实例节点 **默认取值**: 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**: 事件级别。 **约束限制**: 不涉及。 **取值范围**: - CRITICAL：紧急 - MAJOR：重要 - MINOR：一般 - INFO：提示 **默认取值**: 不涉及。
	Level *string `json:"level,omitempty"`

	// **参数解释**: 排序字段。 **约束限制**: 不涉及。 **取值范围**: - planned_execution_time：计划执行时间 - created_time：创建时间 - latest_execution_time：最晚执行时间 **默认取值**: 不涉及。
	SortField *string `json:"sort_field,omitempty"`

	// **参数解释**: 排序顺序。 **约束限制**: 不涉及。 **取值范围**: - DESC：降序 - ASC：升序 **默认取值**: DESC。
	Order *string `json:"order,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。 **约束限制**: 必须为数字，不能为负数。 **取值范围**: 不涉及。 **默认取值**: 0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 查询记录数。 **约束限制**: 不能为负数。 **取值范围**: 最小值为1，最大值为100。 **默认取值**: 10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsRequest struct{}"
	}

	return strings.Join([]string{"ListEventsRequest", string(data)}, " ")
}

type ListEventsRequestXLanguage struct {
	value string
}

type ListEventsRequestXLanguageEnum struct {
	EN_US ListEventsRequestXLanguage
	ZH_CN ListEventsRequestXLanguage
}

func GetListEventsRequestXLanguageEnum() ListEventsRequestXLanguageEnum {
	return ListEventsRequestXLanguageEnum{
		EN_US: ListEventsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListEventsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListEventsRequestXLanguage) Value() string {
	return c.value
}

func (c ListEventsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
