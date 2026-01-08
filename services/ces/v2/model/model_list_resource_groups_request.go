package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceGroupsRequest Request Object
type ListResourceGroupsRequest struct {

	// **参数解释**: 归属企业项目ID。 **约束限制**: 不涉及。 **取值范围**: 只能包含小写字母、数字、“-”、“_”，长度为36个字符。或者为0（代表默认企业项目ID），或者为all_granted_eps（代表所有企业项目ID）。 **默认取值**: 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 资源分组名称，支持模糊查询。 **约束限制** 不涉及。 **取值范围** 包含字母、数字、_、-或汉字，长度为[1,128]个字符。 **默认取值** 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释** 资源分组ID。 **约束限制** 不涉及。 **取值范围** 以\"rg\"开头，后跟22位由字母或数字组成的字符串。 **默认取值** 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释** 分页起始值。 **约束限制** 不涉及。 **取值范围** 在[0,10000]区间内。 **默认取值** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释** 分页查询时每页的条目数。 **约束限制** 不涉及。 **取值范围** 在[1,100]区间内。 **默认取值** 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释** 资源分组添加资源方式，不传代表查询所有资源分组类型。 **约束限制** 不涉及。 **取值范围** - EPS: 表示匹配企业项目 - TAG: 表示匹配标签 - Manual: 表示手动添加 - COMB: 表示组合匹配 - NAME: 表示匹配实例名称 **默认取值** 不涉及。
	Type *ListResourceGroupsRequestType `json:"type,omitempty"`

	// **参数解释** 资源分组健康状态。 **约束限制** 不涉及。 **取值范围** - health: 表示无告警 - unhealthy: 表示告警中 - no_alarm_rule: 表示未设置告警规则 **默认取值** 不涉及。
	Status *ListResourceGroupsRequestStatus `json:"status,omitempty"`
}

func (o ListResourceGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceGroupsRequest", string(data)}, " ")
}

type ListResourceGroupsRequestType struct {
	value string
}

type ListResourceGroupsRequestTypeEnum struct {
	EPS    ListResourceGroupsRequestType
	TAG    ListResourceGroupsRequestType
	MANUAL ListResourceGroupsRequestType
	COMB   ListResourceGroupsRequestType
	NAME   ListResourceGroupsRequestType
}

func GetListResourceGroupsRequestTypeEnum() ListResourceGroupsRequestTypeEnum {
	return ListResourceGroupsRequestTypeEnum{
		EPS: ListResourceGroupsRequestType{
			value: "EPS",
		},
		TAG: ListResourceGroupsRequestType{
			value: "TAG",
		},
		MANUAL: ListResourceGroupsRequestType{
			value: "Manual",
		},
		COMB: ListResourceGroupsRequestType{
			value: "COMB",
		},
		NAME: ListResourceGroupsRequestType{
			value: "NAME",
		},
	}
}

func (c ListResourceGroupsRequestType) Value() string {
	return c.value
}

func (c ListResourceGroupsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceGroupsRequestType) UnmarshalJSON(b []byte) error {
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

type ListResourceGroupsRequestStatus struct {
	value string
}

type ListResourceGroupsRequestStatusEnum struct {
	HEALTH        ListResourceGroupsRequestStatus
	UNHEALTHY     ListResourceGroupsRequestStatus
	NO_ALARM_RULE ListResourceGroupsRequestStatus
}

func GetListResourceGroupsRequestStatusEnum() ListResourceGroupsRequestStatusEnum {
	return ListResourceGroupsRequestStatusEnum{
		HEALTH: ListResourceGroupsRequestStatus{
			value: "health",
		},
		UNHEALTHY: ListResourceGroupsRequestStatus{
			value: "unhealthy",
		},
		NO_ALARM_RULE: ListResourceGroupsRequestStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c ListResourceGroupsRequestStatus) Value() string {
	return c.value
}

func (c ListResourceGroupsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceGroupsRequestStatus) UnmarshalJSON(b []byte) error {
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
