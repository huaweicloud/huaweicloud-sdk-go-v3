package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ApiPolicyMockCreate struct {
	// 返回结果

	ResultContent *string `json:"result_content,omitempty"`
	// 关联的策略组合模式： - ALL：满足全部条件 - ANY：满足任一条件

	EffectMode ApiPolicyMockCreateEffectMode `json:"effect_mode"`
	// 策略后端名称。字符串由中文、英文字母、数字、下划线组成，且只能以中文或英文开头。

	Name string `json:"name"`
	// 后端参数列表

	BackendParams *[]BackendParamBase `json:"backend_params,omitempty"`
	// 策略条件列表

	Conditions []ApiConditionBase `json:"conditions"`
	// 后端自定义认证对象的ID

	AuthorizerId *string `json:"authorizer_id,omitempty"`
}

func (o ApiPolicyMockCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyMockCreate struct{}"
	}

	return strings.Join([]string{"ApiPolicyMockCreate", string(data)}, " ")
}

type ApiPolicyMockCreateEffectMode struct {
	value string
}

type ApiPolicyMockCreateEffectModeEnum struct {
	ALL ApiPolicyMockCreateEffectMode
	ANY ApiPolicyMockCreateEffectMode
}

func GetApiPolicyMockCreateEffectModeEnum() ApiPolicyMockCreateEffectModeEnum {
	return ApiPolicyMockCreateEffectModeEnum{
		ALL: ApiPolicyMockCreateEffectMode{
			value: "ALL",
		},
		ANY: ApiPolicyMockCreateEffectMode{
			value: "ANY",
		},
	}
}

func (c ApiPolicyMockCreateEffectMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyMockCreateEffectMode) UnmarshalJSON(b []byte) error {
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
