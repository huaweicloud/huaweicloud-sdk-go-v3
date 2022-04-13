package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

type ThrottleApiBinding struct {
	// API的发布记录编号

	PublishId *string `json:"publish_id,omitempty"`
	// 策略作用域，取值如下： - 1：整个API - 2： 单个用户 - 3：单个APP  目前只支持1

	Scope *ThrottleApiBindingScope `json:"scope,omitempty"`
	// 流控策略的ID

	StrategyId *string `json:"strategy_id,omitempty"`
	// 绑定时间

	ApplyTime *sdktime.SdkTime `json:"apply_time,omitempty"`
	// 绑定关系的ID

	Id *string `json:"id,omitempty"`
}

func (o ThrottleApiBinding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleApiBinding struct{}"
	}

	return strings.Join([]string{"ThrottleApiBinding", string(data)}, " ")
}

type ThrottleApiBindingScope struct {
	value int32
}

type ThrottleApiBindingScopeEnum struct {
	E_1 ThrottleApiBindingScope
	E_2 ThrottleApiBindingScope
	E_3 ThrottleApiBindingScope
}

func GetThrottleApiBindingScopeEnum() ThrottleApiBindingScopeEnum {
	return ThrottleApiBindingScopeEnum{
		E_1: ThrottleApiBindingScope{
			value: 1,
		}, E_2: ThrottleApiBindingScope{
			value: 2,
		}, E_3: ThrottleApiBindingScope{
			value: 3,
		},
	}
}

func (c ThrottleApiBindingScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThrottleApiBindingScope) UnmarshalJSON(b []byte) error {
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
