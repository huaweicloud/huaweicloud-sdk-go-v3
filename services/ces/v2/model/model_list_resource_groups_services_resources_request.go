package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceGroupsServicesResourcesRequest Request Object
type ListResourceGroupsServicesResourcesRequest struct {

	// **参数解释**： 分页偏移量 **约束限制**： 不涉及 **取值范围**： 整数，[0,10000] **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页大小。 **约束限制**： 不涉及。 **取值范围**： 大小为1-100。 **默认取值**： 不涉及。
	Limit *string `json:"limit,omitempty"`

	// **参数解释** 资源分组ID。 **约束限制** 不涉及 **取值范围** 以\"rg\"开头，后面跟着22个字母或数字 **默认取值** 不涉及
	GroupId string `json:"group_id"`

	// **参数解释** 服务类别，如SYS.ECS **约束限制** 不涉及 **取值范围** 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度在 [3,32]个字符之间 **默认取值** 不涉及
	Service string `json:"service"`

	// **参数解释**： 资源维度。 **约束限制**： 不涉及。 **取值范围**： 多维度用\",\"分割，只能包含0-9、a-z、A-Z、_、-、#、/、(、），每个维度的最大长度为32。字符串总长度最小为1，最大为131。 **默认取值**： 不涉及。
	DimName *string `json:"dim_name,omitempty"`

	// **参数解释** 告警规则按状态信息进行过滤。 **约束限制**： 不涉及。 **取值范围** 枚举值。 - health: 已设置告警规则且无告警触发的资源 - unhealthy: 已设置告警规则且有告警触发的资源 - no_alarm_rule: 未设置告警规则的资源 **默认取值**： 不涉及。
	Status *ListResourceGroupsServicesResourcesRequestStatus `json:"status,omitempty"`

	// **参数描述**： 资源维度值，不支持模糊匹配，但是多维度资源可以只指定一个维度值 **约束限制**： 不涉及。  **取值范围**： 字符串长度范围[1,1027] **默认取值**： 不涉及。
	DimValue *string `json:"dim_value,omitempty"`

	// **参数描述**： 资源的标签信息，格式：\"[key]\":\"[value]\"，样例参考：\"ssss\":\"1111\" **约束限制**： 不涉及。  **取值范围**： 字符串长度范围[0,500] **默认取值**： 不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释** 企业项目ID。 **约束限制** 不涉及。 **取值范围** 由数字、字母和-组成，字符串长度范围[1,128] **默认取值** 不涉及。
	ExtendRelationId *string `json:"extend_relation_id,omitempty"`

	// **参数解释**： 资源所属的云产品名称，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"。 **约束限制**： 不涉及。 **取值范围**： 长度为[0,128]个字符。        **默认取值**： 不涉及。
	ProductName *string `json:"product_name,omitempty"`

	// **参数解释** 资源名称 **约束限制** 不涉及 **取值范围** 长度[1,128]个字符 **默认取值** 不涉及
	ResourceName *string `json:"resource_name,omitempty"`

	// **参数解释** 按事件告警状态信息进行过滤。 **约束限制**： 不涉及。 **取值范围** 枚举值。 - health: 已设置事件告警规则且无事件告警触发的资源 - unhealthy: 已设置事件告警规则且有事件告警触发的资源 - no_alarm_rule: 未设置事件告警规则的资源 **默认取值**： 不涉及。
	EventStatus *ListResourceGroupsServicesResourcesRequestEventStatus `json:"event_status,omitempty"`
}

func (o ListResourceGroupsServicesResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupsServicesResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListResourceGroupsServicesResourcesRequest", string(data)}, " ")
}

type ListResourceGroupsServicesResourcesRequestStatus struct {
	value string
}

type ListResourceGroupsServicesResourcesRequestStatusEnum struct {
	HEALTH        ListResourceGroupsServicesResourcesRequestStatus
	UNHEALTHY     ListResourceGroupsServicesResourcesRequestStatus
	NO_ALARM_RULE ListResourceGroupsServicesResourcesRequestStatus
}

func GetListResourceGroupsServicesResourcesRequestStatusEnum() ListResourceGroupsServicesResourcesRequestStatusEnum {
	return ListResourceGroupsServicesResourcesRequestStatusEnum{
		HEALTH: ListResourceGroupsServicesResourcesRequestStatus{
			value: "health",
		},
		UNHEALTHY: ListResourceGroupsServicesResourcesRequestStatus{
			value: "unhealthy",
		},
		NO_ALARM_RULE: ListResourceGroupsServicesResourcesRequestStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c ListResourceGroupsServicesResourcesRequestStatus) Value() string {
	return c.value
}

func (c ListResourceGroupsServicesResourcesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceGroupsServicesResourcesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListResourceGroupsServicesResourcesRequestEventStatus struct {
	value string
}

type ListResourceGroupsServicesResourcesRequestEventStatusEnum struct {
	HEALTH        ListResourceGroupsServicesResourcesRequestEventStatus
	UNHEALTHY     ListResourceGroupsServicesResourcesRequestEventStatus
	NO_ALARM_RULE ListResourceGroupsServicesResourcesRequestEventStatus
}

func GetListResourceGroupsServicesResourcesRequestEventStatusEnum() ListResourceGroupsServicesResourcesRequestEventStatusEnum {
	return ListResourceGroupsServicesResourcesRequestEventStatusEnum{
		HEALTH: ListResourceGroupsServicesResourcesRequestEventStatus{
			value: "health",
		},
		UNHEALTHY: ListResourceGroupsServicesResourcesRequestEventStatus{
			value: "unhealthy",
		},
		NO_ALARM_RULE: ListResourceGroupsServicesResourcesRequestEventStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c ListResourceGroupsServicesResourcesRequestEventStatus) Value() string {
	return c.value
}

func (c ListResourceGroupsServicesResourcesRequestEventStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceGroupsServicesResourcesRequestEventStatus) UnmarshalJSON(b []byte) error {
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
