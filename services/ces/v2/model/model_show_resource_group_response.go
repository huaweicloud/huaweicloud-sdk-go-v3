package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowResourceGroupResponse Response Object
type ShowResourceGroupResponse struct {

	// 资源分组的名称
	GroupName *string `json:"group_name,omitempty"`

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId *string `json:"group_id,omitempty"`

	// 资源分组的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 资源分组归属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源添加/匹配方式，取值只能为EPS（匹配企业项目）,TAG（匹配标签）,NAME（匹配实例名称）, COMB（组合匹配）,Manual（手动添加）
	Type *ShowResourceGroupResponseType `json:"type,omitempty"`

	// 该资源分组内包含的资源来源的企业项目ID，type为EPS时必传
	AssociationEpIds *[]string `json:"association_ep_ids,omitempty"`

	// 当资源匹配规则为匹配标签时,所指定的标签规则
	Tags *[]ResourceGroupTagRelation `json:"tags,omitempty"`

	// 实例名称匹配参数
	Instances *[]Instance `json:"instances,omitempty"`

	CombRelation *CombRelation `json:"comb_relation,omitempty"`

	// 当资源匹配规则为匹配企业项目时，指定的企业项目列表
	RelatedEpIds *[]string `json:"related_ep_ids,omitempty"`

	// 匹配企业项目或匹配标签参数
	EnterpriseProjectIdAndTags *[]EnterpriseProjectIdAndTags `json:"enterprise_project_id_and_tags,omitempty"`

	// 指标告警状态，取值为health（告警中）、unhealthy（已触发）、no_alarm_rule（未设置告警规则）
	Status *ShowResourceGroupResponseStatus `json:"status,omitempty"`

	// 事件告警状态，取值为health（告警中）、unhealthy（已触发）、no_alarm_rule（未设置告警规则）
	EventStatus *ShowResourceGroupResponseEventStatus `json:"event_status,omitempty"`

	ResourceStatistics *OneResourceGroupRespResourceStatistics `json:"resource_statistics,omitempty"`

	// dimension: 子维度,product: 云产品
	ResourceLevel *ShowResourceGroupResponseResourceLevel `json:"resource_level,omitempty"`

	// 创建资源层级为云产品时的云产品的取值，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"。多个云产品则用“;”隔开，如\"SERVICE.BMS,instance_id;SYS.ECS,instance_id\"。
	ProductNames *string `json:"product_names,omitempty"`

	// 每个企业项目关联的资源状态
	EpResourceStatistics *[]EpResourceStatistics `json:"ep_resource_statistics,omitempty"`
	HttpStatusCode       int                     `json:"-"`
}

func (o ShowResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceGroupResponse", string(data)}, " ")
}

type ShowResourceGroupResponseType struct {
	value string
}

type ShowResourceGroupResponseTypeEnum struct {
	EPS    ShowResourceGroupResponseType
	TAG    ShowResourceGroupResponseType
	NAME   ShowResourceGroupResponseType
	COMB   ShowResourceGroupResponseType
	MANUAL ShowResourceGroupResponseType
}

func GetShowResourceGroupResponseTypeEnum() ShowResourceGroupResponseTypeEnum {
	return ShowResourceGroupResponseTypeEnum{
		EPS: ShowResourceGroupResponseType{
			value: "EPS",
		},
		TAG: ShowResourceGroupResponseType{
			value: "TAG",
		},
		NAME: ShowResourceGroupResponseType{
			value: "NAME",
		},
		COMB: ShowResourceGroupResponseType{
			value: "COMB",
		},
		MANUAL: ShowResourceGroupResponseType{
			value: "Manual",
		},
	}
}

func (c ShowResourceGroupResponseType) Value() string {
	return c.value
}

func (c ShowResourceGroupResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceGroupResponseType) UnmarshalJSON(b []byte) error {
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

type ShowResourceGroupResponseStatus struct {
	value string
}

type ShowResourceGroupResponseStatusEnum struct {
	HEALTH        ShowResourceGroupResponseStatus
	UNHEALTHY     ShowResourceGroupResponseStatus
	NO_ALARM_RULE ShowResourceGroupResponseStatus
}

func GetShowResourceGroupResponseStatusEnum() ShowResourceGroupResponseStatusEnum {
	return ShowResourceGroupResponseStatusEnum{
		HEALTH: ShowResourceGroupResponseStatus{
			value: "health",
		},
		UNHEALTHY: ShowResourceGroupResponseStatus{
			value: "unhealthy",
		},
		NO_ALARM_RULE: ShowResourceGroupResponseStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c ShowResourceGroupResponseStatus) Value() string {
	return c.value
}

func (c ShowResourceGroupResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceGroupResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowResourceGroupResponseEventStatus struct {
	value string
}

type ShowResourceGroupResponseEventStatusEnum struct {
	HEALTH        ShowResourceGroupResponseEventStatus
	UNHEALTHY     ShowResourceGroupResponseEventStatus
	NO_ALARM_RULE ShowResourceGroupResponseEventStatus
}

func GetShowResourceGroupResponseEventStatusEnum() ShowResourceGroupResponseEventStatusEnum {
	return ShowResourceGroupResponseEventStatusEnum{
		HEALTH: ShowResourceGroupResponseEventStatus{
			value: "health",
		},
		UNHEALTHY: ShowResourceGroupResponseEventStatus{
			value: "unhealthy",
		},
		NO_ALARM_RULE: ShowResourceGroupResponseEventStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c ShowResourceGroupResponseEventStatus) Value() string {
	return c.value
}

func (c ShowResourceGroupResponseEventStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceGroupResponseEventStatus) UnmarshalJSON(b []byte) error {
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

type ShowResourceGroupResponseResourceLevel struct {
	value string
}

type ShowResourceGroupResponseResourceLevelEnum struct {
	DIMENSION ShowResourceGroupResponseResourceLevel
	PRODUCT   ShowResourceGroupResponseResourceLevel
}

func GetShowResourceGroupResponseResourceLevelEnum() ShowResourceGroupResponseResourceLevelEnum {
	return ShowResourceGroupResponseResourceLevelEnum{
		DIMENSION: ShowResourceGroupResponseResourceLevel{
			value: "dimension",
		},
		PRODUCT: ShowResourceGroupResponseResourceLevel{
			value: "product",
		},
	}
}

func (c ShowResourceGroupResponseResourceLevel) Value() string {
	return c.value
}

func (c ShowResourceGroupResponseResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceGroupResponseResourceLevel) UnmarshalJSON(b []byte) error {
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
