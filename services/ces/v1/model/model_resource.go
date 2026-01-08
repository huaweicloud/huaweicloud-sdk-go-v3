package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Resource **参数解释**： 资源信息
type Resource struct {

	// **参数解释**： 告警规则的ID或者资源分组ID。 **取值范围**： 不涉及
	RelationId *string `json:"relation_id,omitempty"`

	// **参数解释**： 服务指标命名空间。 **取值范围**： 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释** 指标维度信息
	Dimensions *[]DimensionResp `json:"dimensions,omitempty"`

	// **参数解释** 资源健康状态 **取值范围** - health: 表示无告警 - unhealth: 表示告警中 - no_alarm_rule: 表示未设置告警规则
	Status *ResourceStatus `json:"status,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}

type ResourceStatus struct {
	value string
}

type ResourceStatusEnum struct {
	HEALTH        ResourceStatus
	UNHEALTH      ResourceStatus
	NO_ALARM_RULE ResourceStatus
}

func GetResourceStatusEnum() ResourceStatusEnum {
	return ResourceStatusEnum{
		HEALTH: ResourceStatus{
			value: "health",
		},
		UNHEALTH: ResourceStatus{
			value: "unhealth",
		},
		NO_ALARM_RULE: ResourceStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c ResourceStatus) Value() string {
	return c.value
}

func (c ResourceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceStatus) UnmarshalJSON(b []byte) error {
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
