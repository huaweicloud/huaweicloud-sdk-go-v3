package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type OneResourceGroupResp struct {

	// **参数解释** 资源分组的名称。 **取值范围** 只能为字母、数字、汉字、-或_，长度为[1,128]个字符。
	GroupName string `json:"group_name"`

	// **参数解释** 资源分组ID。 **取值范围** 以rg开头，后跟22位由字母或数字组成的字符串。
	GroupId string `json:"group_id"`

	// **参数解释** 资源分组的创建时间，UNIX时间戳，单位毫秒；如：1603819753000。
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// **参数解释** 资源分组的修改时间，UNIX时间戳，单位毫秒；如：1603819753000。
	UpdateTime *sdktime.SdkTime `json:"update_time"`

	// **参数解释** 资源分组归属企业项目ID。 **取值范围** 只能包含小写字母、数字或-，长度为36个字符。或者为0，代表默认企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// **参数解释** 资源分组添加资源方式。 **取值范围** - EPS: 表示匹配企业项目 - TAG: 表示匹配标签 - Manual: 表示手动添加 - COMB: 表示组合匹配 - NAME: 表示匹配实例名称
	Type OneResourceGroupRespType `json:"type"`

	// **参数解释** 资源分组指标告警状态。 **取值范围** - health: 表示无告警 - unhealthy: 表示告警中 - no_alarm_rule: 表示未设置告警规则
	Status *OneResourceGroupRespStatus `json:"status,omitempty"`

	// **参数解释** 资源分组事件告警状态。 **取值范围** - health: 表示无告警 - unhealthy: 表示告警中 - no_alarm_rule: 表示未设置告警规则
	EventStatus *OneResourceGroupRespEventStatus `json:"event_status,omitempty"`

	ResourceStatistics *OneResourceGroupRespResourceStatistics `json:"resource_statistics,omitempty"`

	// **参数解释** 当资源匹配规则为匹配企业项目时，指定的企业项目列表。
	RelatedEpIds *[]string `json:"related_ep_ids,omitempty"`

	// **参数解释** 关联的告警模板列表。
	AssociationAlarmTemplates *[]AssociationAlarmTemplate `json:"association_alarm_templates,omitempty"`

	// **参数解释** 资源层级，资源生效范围。 **取值范围** - product: 云产品 - dimension: 子维度
	ResourceLevel *OneResourceGroupRespResourceLevel `json:"resource_level,omitempty"`
}

func (o OneResourceGroupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneResourceGroupResp struct{}"
	}

	return strings.Join([]string{"OneResourceGroupResp", string(data)}, " ")
}

type OneResourceGroupRespType struct {
	value string
}

type OneResourceGroupRespTypeEnum struct {
	EPS    OneResourceGroupRespType
	TAG    OneResourceGroupRespType
	MANUAL OneResourceGroupRespType
	COMB   OneResourceGroupRespType
	NAME   OneResourceGroupRespType
}

func GetOneResourceGroupRespTypeEnum() OneResourceGroupRespTypeEnum {
	return OneResourceGroupRespTypeEnum{
		EPS: OneResourceGroupRespType{
			value: "EPS",
		},
		TAG: OneResourceGroupRespType{
			value: "TAG",
		},
		MANUAL: OneResourceGroupRespType{
			value: "Manual",
		},
		COMB: OneResourceGroupRespType{
			value: "COMB",
		},
		NAME: OneResourceGroupRespType{
			value: "NAME",
		},
	}
}

func (c OneResourceGroupRespType) Value() string {
	return c.value
}

func (c OneResourceGroupRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OneResourceGroupRespType) UnmarshalJSON(b []byte) error {
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

type OneResourceGroupRespStatus struct {
	value string
}

type OneResourceGroupRespStatusEnum struct {
	HEALTH        OneResourceGroupRespStatus
	UNHEALTHY     OneResourceGroupRespStatus
	NO_ALARM_RULE OneResourceGroupRespStatus
}

func GetOneResourceGroupRespStatusEnum() OneResourceGroupRespStatusEnum {
	return OneResourceGroupRespStatusEnum{
		HEALTH: OneResourceGroupRespStatus{
			value: "health",
		},
		UNHEALTHY: OneResourceGroupRespStatus{
			value: "unhealthy",
		},
		NO_ALARM_RULE: OneResourceGroupRespStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c OneResourceGroupRespStatus) Value() string {
	return c.value
}

func (c OneResourceGroupRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OneResourceGroupRespStatus) UnmarshalJSON(b []byte) error {
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

type OneResourceGroupRespEventStatus struct {
	value string
}

type OneResourceGroupRespEventStatusEnum struct {
	HEALTH        OneResourceGroupRespEventStatus
	UNHEALTHY     OneResourceGroupRespEventStatus
	NO_ALARM_RULE OneResourceGroupRespEventStatus
}

func GetOneResourceGroupRespEventStatusEnum() OneResourceGroupRespEventStatusEnum {
	return OneResourceGroupRespEventStatusEnum{
		HEALTH: OneResourceGroupRespEventStatus{
			value: "health",
		},
		UNHEALTHY: OneResourceGroupRespEventStatus{
			value: "unhealthy",
		},
		NO_ALARM_RULE: OneResourceGroupRespEventStatus{
			value: "no_alarm_rule",
		},
	}
}

func (c OneResourceGroupRespEventStatus) Value() string {
	return c.value
}

func (c OneResourceGroupRespEventStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OneResourceGroupRespEventStatus) UnmarshalJSON(b []byte) error {
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

type OneResourceGroupRespResourceLevel struct {
	value string
}

type OneResourceGroupRespResourceLevelEnum struct {
	PRODUCT   OneResourceGroupRespResourceLevel
	DIMENSION OneResourceGroupRespResourceLevel
}

func GetOneResourceGroupRespResourceLevelEnum() OneResourceGroupRespResourceLevelEnum {
	return OneResourceGroupRespResourceLevelEnum{
		PRODUCT: OneResourceGroupRespResourceLevel{
			value: "product",
		},
		DIMENSION: OneResourceGroupRespResourceLevel{
			value: "dimension",
		},
	}
}

func (c OneResourceGroupRespResourceLevel) Value() string {
	return c.value
}

func (c OneResourceGroupRespResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OneResourceGroupRespResourceLevel) UnmarshalJSON(b []byte) error {
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
