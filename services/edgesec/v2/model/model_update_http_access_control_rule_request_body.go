package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateHttpAccessControlRuleRequestBody struct {

	// 规则名称
	Name string `json:"name"`

	// 规则描述，最长512字符
	Description *string `json:"description,omitempty"`

	// 规则开关状态
	Status *int32 `json:"status,omitempty"`

	// 生效时间
	Start *int64 `json:"start,omitempty"`

	// 失效时间
	Terminal *int64 `json:"terminal,omitempty"`

	// 生效模式
	TimeMode UpdateHttpAccessControlRuleRequestBodyTimeMode `json:"time_mode"`

	// time_mode为period时必传，每日生效时间类型，目前只有day
	PeriodType *UpdateHttpAccessControlRuleRequestBodyPeriodType `json:"period_type,omitempty"`

	// time_mode为period时必传，每日生效时间区间
	TimeRange *[]TimeRangeItem `json:"time_range,omitempty"`

	// time_mode为period时必传，时区，例如：UTC+8
	TimeZone *string `json:"time_zone,omitempty"`

	// 执行该规则的优先级，值越小，优先级越高，值相同时，规则创建时间早，优先级越高。取值范围：1到100。
	Priority int32 `json:"priority"`

	// 命中条件
	Conditions []HttpAccessControlRuleCondition `json:"conditions"`

	Action *HttpRuleAction `json:"action"`
}

func (o UpdateHttpAccessControlRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpAccessControlRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHttpAccessControlRuleRequestBody", string(data)}, " ")
}

type UpdateHttpAccessControlRuleRequestBodyTimeMode struct {
	value string
}

type UpdateHttpAccessControlRuleRequestBodyTimeModeEnum struct {
	PERMANENT UpdateHttpAccessControlRuleRequestBodyTimeMode
	PERIOD    UpdateHttpAccessControlRuleRequestBodyTimeMode
	CUSTOMIZE UpdateHttpAccessControlRuleRequestBodyTimeMode
}

func GetUpdateHttpAccessControlRuleRequestBodyTimeModeEnum() UpdateHttpAccessControlRuleRequestBodyTimeModeEnum {
	return UpdateHttpAccessControlRuleRequestBodyTimeModeEnum{
		PERMANENT: UpdateHttpAccessControlRuleRequestBodyTimeMode{
			value: "permanent",
		},
		PERIOD: UpdateHttpAccessControlRuleRequestBodyTimeMode{
			value: "period",
		},
		CUSTOMIZE: UpdateHttpAccessControlRuleRequestBodyTimeMode{
			value: "customize",
		},
	}
}

func (c UpdateHttpAccessControlRuleRequestBodyTimeMode) Value() string {
	return c.value
}

func (c UpdateHttpAccessControlRuleRequestBodyTimeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHttpAccessControlRuleRequestBodyTimeMode) UnmarshalJSON(b []byte) error {
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

type UpdateHttpAccessControlRuleRequestBodyPeriodType struct {
	value string
}

type UpdateHttpAccessControlRuleRequestBodyPeriodTypeEnum struct {
	DAY UpdateHttpAccessControlRuleRequestBodyPeriodType
}

func GetUpdateHttpAccessControlRuleRequestBodyPeriodTypeEnum() UpdateHttpAccessControlRuleRequestBodyPeriodTypeEnum {
	return UpdateHttpAccessControlRuleRequestBodyPeriodTypeEnum{
		DAY: UpdateHttpAccessControlRuleRequestBodyPeriodType{
			value: "day",
		},
	}
}

func (c UpdateHttpAccessControlRuleRequestBodyPeriodType) Value() string {
	return c.value
}

func (c UpdateHttpAccessControlRuleRequestBodyPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHttpAccessControlRuleRequestBodyPeriodType) UnmarshalJSON(b []byte) error {
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
