package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BaselineV2 基线任务
type BaselineV2 struct {

	// 基线任务ID。
	Id *string `json:"id,omitempty"`

	// 基线任务名称。
	Name *string `json:"name,omitempty"`

	// 版本号。
	Version *int32 `json:"version,omitempty"`

	// 创建时间戳，单位毫秒。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 最后更新时间戳，单位毫秒。
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 基线任务类型。
	Type *BaselineV2Type `json:"type,omitempty"`

	// 责任人用户ID。
	OwnerId *string `json:"owner_id,omitempty"`

	// 责任人用户名称。
	OwnerName *string `json:"owner_name,omitempty"`

	// 责任人租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 责任人租户名称。
	DomainName *string `json:"domain_name,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// DataArts Studio实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 工作空间ID。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 保障作业ID列表。
	SlaTaskIds *[]string `json:"sla_task_ids,omitempty"`

	// 优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 天基线承诺分钟。
	SlaMin *int32 `json:"sla_min,omitempty"`

	// 预警余量。
	Buffer *int32 `json:"buffer,omitempty"`

	// 天基线承诺小时。
	SlaHour *int32 `json:"sla_hour,omitempty"`

	// 天基线预警分钟。
	ExpMin *int32 `json:"exp_min,omitempty"`

	// 天基线预警小时。
	ExpHour *int32 `json:"exp_hour,omitempty"`

	// 小时基线的预警时间配置（JSON格式），key为周期号，value为hh:mm格式。hh的取值范围为[0,47]，mm的取值范围为[0,59]。
	HourExpDetail *string `json:"hour_exp_detail,omitempty"`

	// 小时基线的承诺时间配置（JSON格式），key为周期号，value为hh:mm格式。hh的取值范围为[0,47]，mm的取值范围为[0,59]。
	HourSlaDetail *string `json:"hour_sla_detail,omitempty"`

	// 是否生效。
	Enable *bool `json:"enable,omitempty"`

	// 报警是否打开。
	AlarmEnable *bool `json:"alarm_enable,omitempty"`

	// 基线报警是否打开。
	BaselineAlarmEnable *bool `json:"baseline_alarm_enable,omitempty"`

	// SMN主题列表。
	SmnTopics *[]SmnTopic `json:"smn_topics,omitempty"`

	// 事件告警开启类型。
	EventAlarm *[]BaselineV2EventAlarm `json:"event_alarm,omitempty"`

	// 事件告警SMN主题列表。
	EventSmnTopics *[]SmnTopic `json:"event_smn_topics,omitempty"`

	// 基线签署是否打开。
	SignEnable *bool `json:"sign_enable,omitempty"`

	// 承诺时间周期列表，小时基线时生效。
	Period *[]PeriodSlaTimeV2 `json:"period,omitempty"`
}

func (o BaselineV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaselineV2 struct{}"
	}

	return strings.Join([]string{"BaselineV2", string(data)}, " ")
}

type BaselineV2Type struct {
	value string
}

type BaselineV2TypeEnum struct {
	DAY  BaselineV2Type
	HOUR BaselineV2Type
}

func GetBaselineV2TypeEnum() BaselineV2TypeEnum {
	return BaselineV2TypeEnum{
		DAY: BaselineV2Type{
			value: "DAY",
		},
		HOUR: BaselineV2Type{
			value: "HOUR",
		},
	}
}

func (c BaselineV2Type) Value() string {
	return c.value
}

func (c BaselineV2Type) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaselineV2Type) UnmarshalJSON(b []byte) error {
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

type BaselineV2EventAlarm struct {
	value string
}

type BaselineV2EventAlarmEnum struct {
	ERROR     BaselineV2EventAlarm
	SLOW_DOWN BaselineV2EventAlarm
}

func GetBaselineV2EventAlarmEnum() BaselineV2EventAlarmEnum {
	return BaselineV2EventAlarmEnum{
		ERROR: BaselineV2EventAlarm{
			value: "ERROR",
		},
		SLOW_DOWN: BaselineV2EventAlarm{
			value: "SLOW_DOWN",
		},
	}
}

func (c BaselineV2EventAlarm) Value() string {
	return c.value
}

func (c BaselineV2EventAlarm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaselineV2EventAlarm) UnmarshalJSON(b []byte) error {
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
