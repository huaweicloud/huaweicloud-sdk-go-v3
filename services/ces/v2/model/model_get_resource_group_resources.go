package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GetResourceGroupResources struct {

	// **参数解释** 按事件告警状态信息进行过滤。 **取值范围** 枚举值。 - health: 告警中 - unhealthy: 已触发 - no_alarm_rule: 未设置告警规则
	Status GetResourceGroupResourcesStatus `json:"status"`

	// 资源的维度信息
	Dimensions []ResourceDimension `json:"dimensions"`

	// **参数解释** 资源的tag信息,格式为key/value的json字符串,样例为\"{\\\"sss\\\":\\\"aaa\\\"}\"。 **取值范围** 枚举值。 字符串长度[0, 10240]
	Tags *string `json:"tags,omitempty"`

	// **参数解释** 企业项目ID。 **取值范围** 由数字、字母和-组成，字符串长度范围[1,128]
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 按事件告警状态信息进行过滤。 **取值范围** 枚举值。 - health: 告警中 - unhealthy: 已触发 - no_alarm_rule: 未设置告警规则
	EventStatus *GetResourceGroupResourcesEventStatus `json:"event_status,omitempty"`

	// **参数解释** 资源名称 **取值范围** 长度[0,128]个字符
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o GetResourceGroupResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceGroupResources struct{}"
	}

	return strings.Join([]string{"GetResourceGroupResources", string(data)}, " ")
}

type GetResourceGroupResourcesStatus struct {
	value string
}

type GetResourceGroupResourcesStatusEnum struct {
	HEALTH        GetResourceGroupResourcesStatus
	UNHEALTHY     GetResourceGroupResourcesStatus
	NO_ALARM_RULE GetResourceGroupResourcesStatus
}

func GetGetResourceGroupResourcesStatusEnum() GetResourceGroupResourcesStatusEnum {
	return GetResourceGroupResourcesStatusEnum{
		HEALTH: GetResourceGroupResourcesStatus{
			value: "health",
		},
		UNHEALTHY: GetResourceGroupResourcesStatus{
			value: "unhealthy",
		},
		NO_ALARM_RULE: GetResourceGroupResourcesStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c GetResourceGroupResourcesStatus) Value() string {
	return c.value
}

func (c GetResourceGroupResourcesStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetResourceGroupResourcesStatus) UnmarshalJSON(b []byte) error {
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

type GetResourceGroupResourcesEventStatus struct {
	value string
}

type GetResourceGroupResourcesEventStatusEnum struct {
	HEALTH        GetResourceGroupResourcesEventStatus
	UNHEALTHY     GetResourceGroupResourcesEventStatus
	NO_ALARM_RULE GetResourceGroupResourcesEventStatus
}

func GetGetResourceGroupResourcesEventStatusEnum() GetResourceGroupResourcesEventStatusEnum {
	return GetResourceGroupResourcesEventStatusEnum{
		HEALTH: GetResourceGroupResourcesEventStatus{
			value: "health",
		},
		UNHEALTHY: GetResourceGroupResourcesEventStatus{
			value: "unhealthy",
		},
		NO_ALARM_RULE: GetResourceGroupResourcesEventStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c GetResourceGroupResourcesEventStatus) Value() string {
	return c.value
}

func (c GetResourceGroupResourcesEventStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetResourceGroupResourcesEventStatus) UnmarshalJSON(b []byte) error {
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
