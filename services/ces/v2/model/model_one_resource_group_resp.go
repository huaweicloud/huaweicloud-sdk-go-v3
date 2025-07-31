package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type OneResourceGroupResp struct {

	// 资源分组的名称
	GroupName string `json:"group_name"`

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId string `json:"group_id"`

	// 资源分组的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 资源分组归属企业项目ID
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 资源添加/匹配方式，取值只能为EPS（匹配企业项目）,TAG（匹配标签）,NAME（匹配实例名称）, COMB（组合匹配）,Manual（手动添加）
	Type OneResourceGroupRespType `json:"type"`

	// 指标告警状态，取值为health（告警中）、unhealthy（已触发）、no_alarm_rule（未设置告警规则）
	Status *OneResourceGroupRespStatus `json:"status,omitempty"`

	// 事件告警状态，取值为health（告警中）、unhealthy（已触发）、no_alarm_rule（未设置告警规则）
	EventStatus *OneResourceGroupRespEventStatus `json:"event_status,omitempty"`

	ResourceStatistics *OneResourceGroupRespResourceStatistics `json:"resource_statistics,omitempty"`

	// 当资源匹配规则为匹配企业项目时，指定的企业项目列表
	RelatedEpIds *[]string `json:"related_ep_ids,omitempty"`

	// 关联的告警模板列表
	AssociationAlarmTemplates *[]AssociationAlarmTemplate `json:"association_alarm_templates,omitempty"`
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
	NAME   OneResourceGroupRespType
	COMB   OneResourceGroupRespType
	MANUAL OneResourceGroupRespType
}

func GetOneResourceGroupRespTypeEnum() OneResourceGroupRespTypeEnum {
	return OneResourceGroupRespTypeEnum{
		EPS: OneResourceGroupRespType{
			value: "EPS",
		},
		TAG: OneResourceGroupRespType{
			value: "TAG",
		},
		NAME: OneResourceGroupRespType{
			value: "NAME",
		},
		COMB: OneResourceGroupRespType{
			value: "COMB",
		},
		MANUAL: OneResourceGroupRespType{
			value: "Manual",
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
