package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateReportProfileDto struct {

	// **参数解释**： 模板描述 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 模板名称 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 发送时间，日报和周报需要设置 **约束限制**： 不涉及 **取值范围**： 0-23 **默认取值**： 不涉及
	SendPeriod *int32 `json:"send_period,omitempty"`

	// **参数解释**： 发送星期，周报需要设置 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SendWeekDay *int32 `json:"send_week_day,omitempty"`

	StatisticPeriod *StatisticPeriod `json:"statistic_period,omitempty"`

	// **参数解释**： 启用状态 **约束限制**： 不涉及 **取值范围**： 0 关闭 1 启用 **默认取值**： 不涉及
	Status *UpdateReportProfileDtoStatus `json:"status,omitempty"`

	// **参数解释**： 通知群组 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释**： 通知方式 **约束限制**： 不涉及 **取值范围**： 0 SMN通知方式 1 不需要通知 **默认取值**： 不涉及
	SubscriptionType *UpdateReportProfileDtoSubscriptionType `json:"subscription_type,omitempty"`
}

func (o UpdateReportProfileDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReportProfileDto struct{}"
	}

	return strings.Join([]string{"UpdateReportProfileDto", string(data)}, " ")
}

type UpdateReportProfileDtoStatus struct {
	value int32
}

type UpdateReportProfileDtoStatusEnum struct {
	E_0 UpdateReportProfileDtoStatus
	E_1 UpdateReportProfileDtoStatus
}

func GetUpdateReportProfileDtoStatusEnum() UpdateReportProfileDtoStatusEnum {
	return UpdateReportProfileDtoStatusEnum{
		E_0: UpdateReportProfileDtoStatus{
			value: 0,
		}, E_1: UpdateReportProfileDtoStatus{
			value: 1,
		},
	}
}

func (c UpdateReportProfileDtoStatus) Value() int32 {
	return c.value
}

func (c UpdateReportProfileDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateReportProfileDtoStatus) UnmarshalJSON(b []byte) error {
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

type UpdateReportProfileDtoSubscriptionType struct {
	value int32
}

type UpdateReportProfileDtoSubscriptionTypeEnum struct {
	E_0 UpdateReportProfileDtoSubscriptionType
	E_1 UpdateReportProfileDtoSubscriptionType
}

func GetUpdateReportProfileDtoSubscriptionTypeEnum() UpdateReportProfileDtoSubscriptionTypeEnum {
	return UpdateReportProfileDtoSubscriptionTypeEnum{
		E_0: UpdateReportProfileDtoSubscriptionType{
			value: 0,
		}, E_1: UpdateReportProfileDtoSubscriptionType{
			value: 1,
		},
	}
}

func (c UpdateReportProfileDtoSubscriptionType) Value() int32 {
	return c.value
}

func (c UpdateReportProfileDtoSubscriptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateReportProfileDtoSubscriptionType) UnmarshalJSON(b []byte) error {
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
