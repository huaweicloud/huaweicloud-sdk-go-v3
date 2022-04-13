package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

// Response Object
type CreateRequestThrottlingPolicyV2Response struct {
	// APP流量限制是指一个API在时长之内被每个APP访问的次数上限，该数值不超过用户流量限制值。输入的值不超过2147483647。正整数。

	AppCallLimits *int32 `json:"app_call_limits,omitempty"`
	// 流控策略名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3 ~ 64字符。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// 流控的时间单位

	TimeUnit CreateRequestThrottlingPolicyV2ResponseTimeUnit `json:"time_unit"`
	// 流控策略描述字符长度不超过255。 > 中文字符必须为UTF-8或者unicode编码。

	Remark *string `json:"remark,omitempty"`
	// API流量限制是指时长内一个API能够被访问的次数上限。该值不超过系统默认配额限制，系统默认配额为200tps，用户可根据实际情况修改该系统默认配额。输入的值不超过2147483647。正整数。

	ApiCallLimits int32 `json:"api_call_limits"`
	// 流控策略的类型 - 1：基础，表示绑定到流控策略的单个API流控时间内能够被调用多少次。 - 2：共享，表示绑定到流控策略的所有API流控时间内能够被调用多少次。

	Type *CreateRequestThrottlingPolicyV2ResponseType `json:"type,omitempty"`
	// 是否开启动态流控： - TRUE - FALSE  暂不支持

	EnableAdaptiveControl *string `json:"enable_adaptive_control,omitempty"`
	// [用户流量限制是指一个API在时长之内每一个用户能访问的次数上限，该数值不超过API流量限制值。输入的值不超过2147483647。正整数。](tag:hws,hws_hk,hcs,fcs,g42)[site不支持用户流量限制,输入值为0](tag:Site)

	UserCallLimits *int32 `json:"user_call_limits,omitempty"`
	// 流量控制的时长单位。与“流量限制次数”配合使用，表示单位时间内的API请求次数上限。输入的值不超过2147483647。正整数。

	TimeInterval int32 `json:"time_interval"`
	// 源IP流量限制是指一个API在时长之内被每个IP访问的次数上限，该数值不超过API流量限制值。输入的值不超过2147483647。正整数。

	IpCallLimits *int32 `json:"ip_call_limits,omitempty"`
	// 流控策略的ID

	Id *string `json:"id,omitempty"`
	// 流控绑定的API数量

	BindNum *int32 `json:"bind_num,omitempty"`
	// 是否包含特殊流控配置 - 1：包含 - 2：不包含

	IsIncluSpecialThrottle *CreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottle `json:"is_inclu_special_throttle,omitempty"`
	// 创建时间

	CreateTime     *sdktime.SdkTime `json:"create_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateRequestThrottlingPolicyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRequestThrottlingPolicyV2Response struct{}"
	}

	return strings.Join([]string{"CreateRequestThrottlingPolicyV2Response", string(data)}, " ")
}

type CreateRequestThrottlingPolicyV2ResponseTimeUnit struct {
	value string
}

type CreateRequestThrottlingPolicyV2ResponseTimeUnitEnum struct {
	SECOND CreateRequestThrottlingPolicyV2ResponseTimeUnit
	MINUTE CreateRequestThrottlingPolicyV2ResponseTimeUnit
	HOUR   CreateRequestThrottlingPolicyV2ResponseTimeUnit
	DAY    CreateRequestThrottlingPolicyV2ResponseTimeUnit
}

func GetCreateRequestThrottlingPolicyV2ResponseTimeUnitEnum() CreateRequestThrottlingPolicyV2ResponseTimeUnitEnum {
	return CreateRequestThrottlingPolicyV2ResponseTimeUnitEnum{
		SECOND: CreateRequestThrottlingPolicyV2ResponseTimeUnit{
			value: "SECOND",
		},
		MINUTE: CreateRequestThrottlingPolicyV2ResponseTimeUnit{
			value: "MINUTE",
		},
		HOUR: CreateRequestThrottlingPolicyV2ResponseTimeUnit{
			value: "HOUR",
		},
		DAY: CreateRequestThrottlingPolicyV2ResponseTimeUnit{
			value: "DAY",
		},
	}
}

func (c CreateRequestThrottlingPolicyV2ResponseTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRequestThrottlingPolicyV2ResponseTimeUnit) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateRequestThrottlingPolicyV2ResponseType struct {
	value int32
}

type CreateRequestThrottlingPolicyV2ResponseTypeEnum struct {
	E_1 CreateRequestThrottlingPolicyV2ResponseType
	E_2 CreateRequestThrottlingPolicyV2ResponseType
}

func GetCreateRequestThrottlingPolicyV2ResponseTypeEnum() CreateRequestThrottlingPolicyV2ResponseTypeEnum {
	return CreateRequestThrottlingPolicyV2ResponseTypeEnum{
		E_1: CreateRequestThrottlingPolicyV2ResponseType{
			value: 1,
		}, E_2: CreateRequestThrottlingPolicyV2ResponseType{
			value: 2,
		},
	}
}

func (c CreateRequestThrottlingPolicyV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRequestThrottlingPolicyV2ResponseType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottle struct {
	value int32
}

type CreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottleEnum struct {
	E_1 CreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottle
	E_2 CreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottle
}

func GetCreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottleEnum() CreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottleEnum {
	return CreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottleEnum{
		E_1: CreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottle{
			value: 1,
		}, E_2: CreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottle{
			value: 2,
		},
	}
}

func (c CreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRequestThrottlingPolicyV2ResponseIsIncluSpecialThrottle) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
