package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateHttpAccessControlRuleRequestBody struct {

	// 规则名称
	Name string `json:"name"`

	// 规则描述，最长512字符
	Description *string `json:"description,omitempty"`

	// 精准防护规则生效的起始时间戳（秒）。当time=true，才需要填写该参数。
	Start *int64 `json:"start,omitempty"`

	// 精准防护规则生效的终止时间戳（秒）。当time=true，才需要填写该参数。
	Terminal *int64 `json:"terminal,omitempty"`

	// 生效模式
	TimeMode CreateHttpAccessControlRuleRequestBodyTimeMode `json:"time_mode"`

	// time_mode为period时必传，每日生效时间类型，目前只有day
	PeriodType *CreateHttpAccessControlRuleRequestBodyPeriodType `json:"period_type,omitempty"`

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

func (o CreateHttpAccessControlRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpAccessControlRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHttpAccessControlRuleRequestBody", string(data)}, " ")
}

type CreateHttpAccessControlRuleRequestBodyTimeMode struct {
	value string
}

type CreateHttpAccessControlRuleRequestBodyTimeModeEnum struct {
	PERMANENT CreateHttpAccessControlRuleRequestBodyTimeMode
	PERIOD    CreateHttpAccessControlRuleRequestBodyTimeMode
	CUSTOMIZE CreateHttpAccessControlRuleRequestBodyTimeMode
}

func GetCreateHttpAccessControlRuleRequestBodyTimeModeEnum() CreateHttpAccessControlRuleRequestBodyTimeModeEnum {
	return CreateHttpAccessControlRuleRequestBodyTimeModeEnum{
		PERMANENT: CreateHttpAccessControlRuleRequestBodyTimeMode{
			value: "permanent",
		},
		PERIOD: CreateHttpAccessControlRuleRequestBodyTimeMode{
			value: "period",
		},
		CUSTOMIZE: CreateHttpAccessControlRuleRequestBodyTimeMode{
			value: "customize",
		},
	}
}

func (c CreateHttpAccessControlRuleRequestBodyTimeMode) Value() string {
	return c.value
}

func (c CreateHttpAccessControlRuleRequestBodyTimeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHttpAccessControlRuleRequestBodyTimeMode) UnmarshalJSON(b []byte) error {
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

type CreateHttpAccessControlRuleRequestBodyPeriodType struct {
	value string
}

type CreateHttpAccessControlRuleRequestBodyPeriodTypeEnum struct {
	DAY CreateHttpAccessControlRuleRequestBodyPeriodType
}

func GetCreateHttpAccessControlRuleRequestBodyPeriodTypeEnum() CreateHttpAccessControlRuleRequestBodyPeriodTypeEnum {
	return CreateHttpAccessControlRuleRequestBodyPeriodTypeEnum{
		DAY: CreateHttpAccessControlRuleRequestBodyPeriodType{
			value: "day",
		},
	}
}

func (c CreateHttpAccessControlRuleRequestBodyPeriodType) Value() string {
	return c.value
}

func (c CreateHttpAccessControlRuleRequestBodyPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHttpAccessControlRuleRequestBodyPeriodType) UnmarshalJSON(b []byte) error {
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
