package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateReportProfileDto struct {

	// **参数解释**： 报告类型 **约束限制**： 不涉及 **取值范围**： daily 日报 weekly 周报 custom 自定义报告 **默认取值**： 不涉及
	Category CreateReportProfileDtoCategory `json:"category"`

	// **参数解释**： 描述信息 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 模板名称 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name string `json:"name"`

	// **参数解释**： 发送时间，日报和周报需要设置 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SendPeriod *int32 `json:"send_period,omitempty"`

	// **参数解释**： 发送星期，周报需要设置 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SendWeekDay *int32 `json:"send_week_day,omitempty"`

	StatisticPeriod *StatisticPeriod `json:"statistic_period,omitempty"`

	// **参数解释**： 启用状态 **约束限制**： 不涉及 **取值范围**： 0 关闭 1 启用 **默认取值**： 不涉及
	Status *CreateReportProfileDtoStatus `json:"status,omitempty"`

	// **参数解释**： 通知群组 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	TopicUrn string `json:"topic_urn"`

	// **参数解释**： 通知方式 **约束限制**： 不涉及 **取值范围**： 0 SMN通知方式 1 不需要通知 **默认取值**： 不涉及
	SubscriptionType *CreateReportProfileDtoSubscriptionType `json:"subscription_type,omitempty"`
}

func (o CreateReportProfileDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportProfileDto struct{}"
	}

	return strings.Join([]string{"CreateReportProfileDto", string(data)}, " ")
}

type CreateReportProfileDtoCategory struct {
	value string
}

type CreateReportProfileDtoCategoryEnum struct {
	DAILY  CreateReportProfileDtoCategory
	WEEKLY CreateReportProfileDtoCategory
	CUSTOM CreateReportProfileDtoCategory
}

func GetCreateReportProfileDtoCategoryEnum() CreateReportProfileDtoCategoryEnum {
	return CreateReportProfileDtoCategoryEnum{
		DAILY: CreateReportProfileDtoCategory{
			value: "daily",
		},
		WEEKLY: CreateReportProfileDtoCategory{
			value: "weekly",
		},
		CUSTOM: CreateReportProfileDtoCategory{
			value: "custom",
		},
	}
}

func (c CreateReportProfileDtoCategory) Value() string {
	return c.value
}

func (c CreateReportProfileDtoCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateReportProfileDtoCategory) UnmarshalJSON(b []byte) error {
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

type CreateReportProfileDtoStatus struct {
	value int32
}

type CreateReportProfileDtoStatusEnum struct {
	E_0 CreateReportProfileDtoStatus
	E_1 CreateReportProfileDtoStatus
}

func GetCreateReportProfileDtoStatusEnum() CreateReportProfileDtoStatusEnum {
	return CreateReportProfileDtoStatusEnum{
		E_0: CreateReportProfileDtoStatus{
			value: 0,
		}, E_1: CreateReportProfileDtoStatus{
			value: 1,
		},
	}
}

func (c CreateReportProfileDtoStatus) Value() int32 {
	return c.value
}

func (c CreateReportProfileDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateReportProfileDtoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateReportProfileDtoSubscriptionType struct {
	value int32
}

type CreateReportProfileDtoSubscriptionTypeEnum struct {
	E_0 CreateReportProfileDtoSubscriptionType
	E_1 CreateReportProfileDtoSubscriptionType
}

func GetCreateReportProfileDtoSubscriptionTypeEnum() CreateReportProfileDtoSubscriptionTypeEnum {
	return CreateReportProfileDtoSubscriptionTypeEnum{
		E_0: CreateReportProfileDtoSubscriptionType{
			value: 0,
		}, E_1: CreateReportProfileDtoSubscriptionType{
			value: 1,
		},
	}
}

func (c CreateReportProfileDtoSubscriptionType) Value() int32 {
	return c.value
}

func (c CreateReportProfileDtoSubscriptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateReportProfileDtoSubscriptionType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
