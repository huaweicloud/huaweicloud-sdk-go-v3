package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceGroupsServicesResourcesRequest Request Object
type ListResourceGroupsServicesResourcesRequest struct {

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId string `json:"group_id"`

	// 服务类别，如SYS.ECS
	Service string `json:"service"`

	// 资源维度信息，多个维度按字母序使用逗号分割
	DimName *string `json:"dim_name,omitempty"`

	// 分页查询时每页的条目数，取值[1,100]，默认值为100
	Limit *string `json:"limit,omitempty"`

	// 分页查询时查询的起始位置，表示从第几条数据开始，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 按状态信息进行过滤，取值只能为health（已设置告警规则且无告警触发的资源）、unhealthy（已设置告警规则且有告警触发的资源）、no_alarm_rule（未设置告警规则的资源）
	Status *ListResourceGroupsServicesResourcesRequestStatus `json:"status,omitempty"`

	// 资源维度值，不支持模糊匹配，但是多维度资源可以只指定一个维度值
	DimValue *string `json:"dim_value,omitempty"`

	// 资源的标签信息，格式：\"[key]\":\"[value]\"，样例参考：\"ssss\":\"1111\"
	Tag *string `json:"tag,omitempty"`

	// 企业项目ID
	ExtendRelationId *string `json:"extend_relation_id,omitempty"`

	// 资源分组的云产品，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"
	ProductName *string `json:"product_name,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 按事件告警状态信息进行过滤，取值只能为health（已设置事件告警规则且无事件告警触发的资源）、unhealthy（已设置事件告警规则且有事件告警触发的资源）、no_alarm_rule（未设置事件告警规则的资源）
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
