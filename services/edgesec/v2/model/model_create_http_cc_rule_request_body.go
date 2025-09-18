package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateHttpCcRuleRequestBody struct {

	// 规则名称
	Name string `json:"name"`

	// cc规则优先级，越大优先级越高，默认1，创建必填
	Priority int32 `json:"priority"`

	// 规则描述，最长512字符
	Description *string `json:"description,omitempty"`

	// cc规则防护模式，现在只支持创建高级cc规则防护模式。   - 0：标准，只支持对域名的防护路径做限制。  - 1：高级，支持对路径、IP、Cookie、Header、Params字段做限制。
	Mode int32 `json:"mode"`

	// 所有用户的周期内请求次数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 单个用户的周期内请求次数
	LimitNum int32 `json:"limit_num"`

	// 限速周期
	LimitPeriod int32 `json:"limit_period"`

	// 阻断时长
	LockTime *int32 `json:"lock_time,omitempty"`

	// 限速模式：   - ip：IP限速，根据IP区分单个Web访问者。   - cookie：用户限速，根据Cookie键值区分单个Web访问者   - header：用户限速，根据Header区分单个Web访问者。   - ip_segment_c：根据IP C段区分单个Web访问者。
	TagType CreateHttpCcRuleRequestBodyTagType `json:"tag_type"`

	// 防护模式标签
	TagIndex *string `json:"tag_index,omitempty"`

	TagCondition *HttpCcRuleCondition `json:"tag_condition,omitempty"`

	// 放行次数
	UnlockNum *int32 `json:"unlock_num,omitempty"`

	// 是否聚合域名
	DomainAggregation *bool `json:"domain_aggregation,omitempty"`

	// cc规则防护规则限速条件，当cc防护规则为高级模式（mode参数值为1）时，该参数必填。
	Conditions *[]HttpCcRuleCondition `json:"conditions,omitempty"`

	Action *HttpRuleAction `json:"action"`

	// 生效模式
	TimeMode CreateHttpCcRuleRequestBodyTimeMode `json:"time_mode"`

	// customize生效时间区间开始
	Start *int64 `json:"start,omitempty"`

	// customize生效时间区间结束
	Terminal *int64 `json:"terminal,omitempty"`

	// period每日生效时间类型，目前只有day
	PeriodType *CreateHttpCcRuleRequestBodyPeriodType `json:"period_type,omitempty"`

	// period每日生效时间区间
	TimeRange *[]TimeRangeItem `json:"time_range,omitempty"`
}

func (o CreateHttpCcRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpCcRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHttpCcRuleRequestBody", string(data)}, " ")
}

type CreateHttpCcRuleRequestBodyTagType struct {
	value string
}

type CreateHttpCcRuleRequestBodyTagTypeEnum struct {
	IP           CreateHttpCcRuleRequestBodyTagType
	COOKIE       CreateHttpCcRuleRequestBodyTagType
	HEADER       CreateHttpCcRuleRequestBodyTagType
	IP_SEGMENT_C CreateHttpCcRuleRequestBodyTagType
}

func GetCreateHttpCcRuleRequestBodyTagTypeEnum() CreateHttpCcRuleRequestBodyTagTypeEnum {
	return CreateHttpCcRuleRequestBodyTagTypeEnum{
		IP: CreateHttpCcRuleRequestBodyTagType{
			value: "ip",
		},
		COOKIE: CreateHttpCcRuleRequestBodyTagType{
			value: "cookie",
		},
		HEADER: CreateHttpCcRuleRequestBodyTagType{
			value: "header",
		},
		IP_SEGMENT_C: CreateHttpCcRuleRequestBodyTagType{
			value: "ip_segment_c",
		},
	}
}

func (c CreateHttpCcRuleRequestBodyTagType) Value() string {
	return c.value
}

func (c CreateHttpCcRuleRequestBodyTagType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHttpCcRuleRequestBodyTagType) UnmarshalJSON(b []byte) error {
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

type CreateHttpCcRuleRequestBodyTimeMode struct {
	value string
}

type CreateHttpCcRuleRequestBodyTimeModeEnum struct {
	PERMANENT CreateHttpCcRuleRequestBodyTimeMode
	PERIOD    CreateHttpCcRuleRequestBodyTimeMode
	CUSTOMIZE CreateHttpCcRuleRequestBodyTimeMode
}

func GetCreateHttpCcRuleRequestBodyTimeModeEnum() CreateHttpCcRuleRequestBodyTimeModeEnum {
	return CreateHttpCcRuleRequestBodyTimeModeEnum{
		PERMANENT: CreateHttpCcRuleRequestBodyTimeMode{
			value: "permanent",
		},
		PERIOD: CreateHttpCcRuleRequestBodyTimeMode{
			value: "period",
		},
		CUSTOMIZE: CreateHttpCcRuleRequestBodyTimeMode{
			value: "customize",
		},
	}
}

func (c CreateHttpCcRuleRequestBodyTimeMode) Value() string {
	return c.value
}

func (c CreateHttpCcRuleRequestBodyTimeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHttpCcRuleRequestBodyTimeMode) UnmarshalJSON(b []byte) error {
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

type CreateHttpCcRuleRequestBodyPeriodType struct {
	value string
}

type CreateHttpCcRuleRequestBodyPeriodTypeEnum struct {
	DAY CreateHttpCcRuleRequestBodyPeriodType
}

func GetCreateHttpCcRuleRequestBodyPeriodTypeEnum() CreateHttpCcRuleRequestBodyPeriodTypeEnum {
	return CreateHttpCcRuleRequestBodyPeriodTypeEnum{
		DAY: CreateHttpCcRuleRequestBodyPeriodType{
			value: "day",
		},
	}
}

func (c CreateHttpCcRuleRequestBodyPeriodType) Value() string {
	return c.value
}

func (c CreateHttpCcRuleRequestBodyPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHttpCcRuleRequestBodyPeriodType) UnmarshalJSON(b []byte) error {
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
