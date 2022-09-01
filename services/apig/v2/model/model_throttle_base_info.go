package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ThrottleBaseInfo struct {

	// APP流量限制是指一个API在时长之内被每个APP访问的次数上限，该数值不超过用户流量限制值。输入的值不超过2147483647。正整数。
	AppCallLimits *int32 `json:"app_call_limits,omitempty" xml:"app_call_limits"`

	// 流控策略名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3 ~ 64字符。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name" xml:"name"`

	// 流控的时间单位
	TimeUnit ThrottleBaseInfoTimeUnit `json:"time_unit" xml:"time_unit"`

	// 流控策略描述字符长度不超过255。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// API流量限制是指时长内一个API能够被访问的次数上限。该值不超过系统默认配额限制，系统默认配额为200tps，用户可根据实际情况修改该系统默认配额。输入的值不超过2147483647。正整数。
	ApiCallLimits int32 `json:"api_call_limits" xml:"api_call_limits"`

	// 流控策略的类型 - 1：基础，表示绑定到流控策略的单个API流控时间内能够被调用多少次。 - 2：共享，表示绑定到流控策略的所有API流控时间内能够被调用多少次。
	Type *ThrottleBaseInfoType `json:"type,omitempty" xml:"type"`

	// 是否开启动态流控： - TRUE - FALSE  暂不支持
	EnableAdaptiveControl *string `json:"enable_adaptive_control,omitempty" xml:"enable_adaptive_control"`

	// [用户流量限制是指一个API在时长之内每一个用户能访问的次数上限，该数值不超过API流量限制值。输入的值不超过2147483647。正整数。](tag:hws,hws_hk,hcs,fcs,g42)[site不支持用户流量限制,输入值为0](tag:Site)
	UserCallLimits *int32 `json:"user_call_limits,omitempty" xml:"user_call_limits"`

	// 流量控制的时长单位。与“流量限制次数”配合使用，表示单位时间内的API请求次数上限。输入的值不超过2147483647。正整数。
	TimeInterval int32 `json:"time_interval" xml:"time_interval"`

	// 源IP流量限制是指一个API在时长之内被每个IP访问的次数上限，该数值不超过API流量限制值。输入的值不超过2147483647。正整数。
	IpCallLimits *int32 `json:"ip_call_limits,omitempty" xml:"ip_call_limits"`
}

func (o ThrottleBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleBaseInfo struct{}"
	}

	return strings.Join([]string{"ThrottleBaseInfo", string(data)}, " ")
}

type ThrottleBaseInfoTimeUnit struct {
	value string
}

type ThrottleBaseInfoTimeUnitEnum struct {
	SECOND ThrottleBaseInfoTimeUnit
	MINUTE ThrottleBaseInfoTimeUnit
	HOUR   ThrottleBaseInfoTimeUnit
	DAY    ThrottleBaseInfoTimeUnit
}

func GetThrottleBaseInfoTimeUnitEnum() ThrottleBaseInfoTimeUnitEnum {
	return ThrottleBaseInfoTimeUnitEnum{
		SECOND: ThrottleBaseInfoTimeUnit{
			value: "SECOND",
		},
		MINUTE: ThrottleBaseInfoTimeUnit{
			value: "MINUTE",
		},
		HOUR: ThrottleBaseInfoTimeUnit{
			value: "HOUR",
		},
		DAY: ThrottleBaseInfoTimeUnit{
			value: "DAY",
		},
	}
}

func (c ThrottleBaseInfoTimeUnit) Value() string {
	return c.value
}

func (c ThrottleBaseInfoTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThrottleBaseInfoTimeUnit) UnmarshalJSON(b []byte) error {
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

type ThrottleBaseInfoType struct {
	value int32
}

type ThrottleBaseInfoTypeEnum struct {
	E_1 ThrottleBaseInfoType
	E_2 ThrottleBaseInfoType
}

func GetThrottleBaseInfoTypeEnum() ThrottleBaseInfoTypeEnum {
	return ThrottleBaseInfoTypeEnum{
		E_1: ThrottleBaseInfoType{
			value: 1,
		}, E_2: ThrottleBaseInfoType{
			value: 2,
		},
	}
}

func (c ThrottleBaseInfoType) Value() int32 {
	return c.value
}

func (c ThrottleBaseInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThrottleBaseInfoType) UnmarshalJSON(b []byte) error {
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
