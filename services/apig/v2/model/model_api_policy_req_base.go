package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPolicyReqBase struct {

	// 关联的策略组合模式： - ALL：满足全部条件 - ANY：满足任一条件
	EffectMode ApiPolicyReqBaseEffectMode `json:"effect_mode"`

	// 策略后端名称。字符串由中文、英文字母、数字、下划线组成，且只能以中文或英文开头。
	Name string `json:"name"`

	// 后端参数列表
	BackendParams *[]BackendParamBase `json:"backend_params,omitempty"`

	// 策略条件列表
	Conditions []ApiConditionBase `json:"conditions"`

	// 后端自定义认证对象的ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`
}

func (o ApiPolicyReqBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyReqBase struct{}"
	}

	return strings.Join([]string{"ApiPolicyReqBase", string(data)}, " ")
}

type ApiPolicyReqBaseEffectMode struct {
	value string
}

type ApiPolicyReqBaseEffectModeEnum struct {
	ALL ApiPolicyReqBaseEffectMode
	ANY ApiPolicyReqBaseEffectMode
}

func GetApiPolicyReqBaseEffectModeEnum() ApiPolicyReqBaseEffectModeEnum {
	return ApiPolicyReqBaseEffectModeEnum{
		ALL: ApiPolicyReqBaseEffectMode{
			value: "ALL",
		},
		ANY: ApiPolicyReqBaseEffectMode{
			value: "ANY",
		},
	}
}

func (c ApiPolicyReqBaseEffectMode) Value() string {
	return c.value
}

func (c ApiPolicyReqBaseEffectMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyReqBaseEffectMode) UnmarshalJSON(b []byte) error {
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
