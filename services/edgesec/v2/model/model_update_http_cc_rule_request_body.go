package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateHttpCcRuleRequestBody struct {

	// 规则名称
	Name string `json:"name"`

	// cc规则优先级，越大优先级越高，默认1
	Priority int32 `json:"priority"`

	// 规则描述，最长512字符
	Description *string `json:"description,omitempty"`

	// 规则开关状态
	Status *int32 `json:"status,omitempty"`

	// cc规则防护模式，现在只支持创建高级cc规则防护模式。   - 0：标准，只支持对域名的防护路径做限制。  - 1：高级，支持对路径、IP、Cookie、Header、Params字段做限制。
	Mode int32 `json:"mode"`

	// 所有用户的周期内请求次数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 单个用户的周期内请求次数
	LimitNum int32 `json:"limit_num"`

	// 限速周期
	LimitPeriod int32 `json:"limit_period"`

	// 锁定时长
	LockTime *int32 `json:"lock_time,omitempty"`

	// 限速模式：   - ip：IP限速，根据IP区分单个Web访问者。   - cookie：用户限速，根据Cookie键值区分单个Web访问者。   - header：用户限速，根据Header区分单个Web访问者。   - other：根据Referer（自定义请求访问的来源）字段区分单个Web访问者。   - policy: 策略限速   - domain: 域名限速     - url: url限速
	TagType UpdateHttpCcRuleRequestBodyTagType `json:"tag_type"`

	// 用户标识，当限速模式为用户限速(cookie或header)时，需要传该参数。   - 选择cookie时，设置cookie字段名，即用户需要根据网站实际情况配置唯一可识别Web访问者的cookie中的某属性变量名。用户标识的cookie，不支持正则，必须完全匹配。例如：如果网站使用cookie中的某个字段name唯一标识用户，那么可以选取name字段来区分Web访问者。   - 选择header时，设置需要防护的自定义HTTP首部，即用户需要根据网站实际情况配置可识别Web访问者的HTTP首部。
	TagIndex *string `json:"tag_index,omitempty"`

	TagCondition *HttpCcRuleCondition `json:"tag_condition,omitempty"`

	// 放行次数
	UnlockNum *int32 `json:"unlock_num,omitempty"`

	// 是否聚合域名
	DomainAggregation *bool `json:"domain_aggregation,omitempty"`

	// 区分大小写，默认不区分false，统一存放小写
	ValueCase *bool `json:"value_case,omitempty"`

	// 限速条件
	Conditions []HttpCcRuleCondition `json:"conditions"`

	Action *HttpRuleAction `json:"action"`

	// 生效模式
	TimeMode *UpdateHttpCcRuleRequestBodyTimeMode `json:"time_mode,omitempty"`

	// customize生效时间区间开始
	Start *int64 `json:"start,omitempty"`

	// customize生效时间区间结束
	Terminal *int64 `json:"terminal,omitempty"`

	// period每日生效时间类型，目前只有day
	PeriodType *UpdateHttpCcRuleRequestBodyPeriodType `json:"period_type,omitempty"`

	// period每日生效时间区间
	TimeRange *[]TimeRangeItem `json:"time_range,omitempty"`
}

func (o UpdateHttpCcRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpCcRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHttpCcRuleRequestBody", string(data)}, " ")
}

type UpdateHttpCcRuleRequestBodyTagType struct {
	value string
}

type UpdateHttpCcRuleRequestBodyTagTypeEnum struct {
	IP     UpdateHttpCcRuleRequestBodyTagType
	COOKIE UpdateHttpCcRuleRequestBodyTagType
}

func GetUpdateHttpCcRuleRequestBodyTagTypeEnum() UpdateHttpCcRuleRequestBodyTagTypeEnum {
	return UpdateHttpCcRuleRequestBodyTagTypeEnum{
		IP: UpdateHttpCcRuleRequestBodyTagType{
			value: "ip",
		},
		COOKIE: UpdateHttpCcRuleRequestBodyTagType{
			value: "cookie",
		},
	}
}

func (c UpdateHttpCcRuleRequestBodyTagType) Value() string {
	return c.value
}

func (c UpdateHttpCcRuleRequestBodyTagType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHttpCcRuleRequestBodyTagType) UnmarshalJSON(b []byte) error {
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

type UpdateHttpCcRuleRequestBodyTimeMode struct {
	value string
}

type UpdateHttpCcRuleRequestBodyTimeModeEnum struct {
	PERMANENT UpdateHttpCcRuleRequestBodyTimeMode
	PERIOD    UpdateHttpCcRuleRequestBodyTimeMode
	CUSTOMIZE UpdateHttpCcRuleRequestBodyTimeMode
}

func GetUpdateHttpCcRuleRequestBodyTimeModeEnum() UpdateHttpCcRuleRequestBodyTimeModeEnum {
	return UpdateHttpCcRuleRequestBodyTimeModeEnum{
		PERMANENT: UpdateHttpCcRuleRequestBodyTimeMode{
			value: "permanent",
		},
		PERIOD: UpdateHttpCcRuleRequestBodyTimeMode{
			value: "period",
		},
		CUSTOMIZE: UpdateHttpCcRuleRequestBodyTimeMode{
			value: "customize",
		},
	}
}

func (c UpdateHttpCcRuleRequestBodyTimeMode) Value() string {
	return c.value
}

func (c UpdateHttpCcRuleRequestBodyTimeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHttpCcRuleRequestBodyTimeMode) UnmarshalJSON(b []byte) error {
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

type UpdateHttpCcRuleRequestBodyPeriodType struct {
	value string
}

type UpdateHttpCcRuleRequestBodyPeriodTypeEnum struct {
	DAY UpdateHttpCcRuleRequestBodyPeriodType
}

func GetUpdateHttpCcRuleRequestBodyPeriodTypeEnum() UpdateHttpCcRuleRequestBodyPeriodTypeEnum {
	return UpdateHttpCcRuleRequestBodyPeriodTypeEnum{
		DAY: UpdateHttpCcRuleRequestBodyPeriodType{
			value: "day",
		},
	}
}

func (c UpdateHttpCcRuleRequestBodyPeriodType) Value() string {
	return c.value
}

func (c UpdateHttpCcRuleRequestBodyPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHttpCcRuleRequestBodyPeriodType) UnmarshalJSON(b []byte) error {
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
