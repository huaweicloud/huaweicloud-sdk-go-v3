package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceGroupRequest Request Object
type ListResourceGroupRequest struct {

	// **参数解释** 资源分组名称。 **约束限制** 不涉及。 **取值范围** 包含字母、数字、_、-或汉字，长度为[1,128]个字符。 **默认取值** 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释** 资源分组ID。 **约束限制** 不涉及。 **取值范围** 以\"rg\"开头，后面跟着22个字母或数字。 **默认取值** 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释** 资源分组健康状态。 **约束限制** 不涉及。 **取值范围** - health: 表示健康 - unhealth: 表示不健康 - no_alarm_rule: 表示未配置告警规则 **默认取值** 不涉及。
	Status *ListResourceGroupRequestStatus `json:"status,omitempty"`

	// **参数解释** 分页起始值。 **约束限制** 不涉及。 **取值范围** 在[0,9999999]区间内。 **默认取值** 0
	Start *int32 `json:"start,omitempty"`

	// **参数解释** 单次查询的条数限制。 **约束限制** 不涉及。 **取值范围** 在[1,100]区间内。 **默认取值** 100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"ListResourceGroupRequest", string(data)}, " ")
}

type ListResourceGroupRequestStatus struct {
	value string
}

type ListResourceGroupRequestStatusEnum struct {
	HEALTH        ListResourceGroupRequestStatus
	UNHEALTH      ListResourceGroupRequestStatus
	NO_ALARM_RULE ListResourceGroupRequestStatus
}

func GetListResourceGroupRequestStatusEnum() ListResourceGroupRequestStatusEnum {
	return ListResourceGroupRequestStatusEnum{
		HEALTH: ListResourceGroupRequestStatus{
			value: "health",
		},
		UNHEALTH: ListResourceGroupRequestStatus{
			value: "unhealth",
		},
		NO_ALARM_RULE: ListResourceGroupRequestStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c ListResourceGroupRequestStatus) Value() string {
	return c.value
}

func (c ListResourceGroupRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceGroupRequestStatus) UnmarshalJSON(b []byte) error {
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
