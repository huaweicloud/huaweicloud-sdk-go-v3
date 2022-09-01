package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPolicyRespBase struct {

	// 编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 策略后端名称。字符串由中文、英文字母、数字、下划线组成，且只能以中文或英文开头。
	Name string `json:"name" xml:"name"`

	// 策略条件列表
	Conditions []ConditionResp `json:"conditions" xml:"conditions"`

	// 后端参数列表
	BackendParams *[]BackendParam `json:"backend_params,omitempty" xml:"backend_params"`

	// 关联的策略组合模式： - ALL：满足全部条件 - ANY：满足任一条件
	EffectMode ApiPolicyRespBaseEffectMode `json:"effect_mode" xml:"effect_mode"`

	// 后端自定义认证对象的ID
	AuthorizerId *string `json:"authorizer_id,omitempty" xml:"authorizer_id"`
}

func (o ApiPolicyRespBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyRespBase struct{}"
	}

	return strings.Join([]string{"ApiPolicyRespBase", string(data)}, " ")
}

type ApiPolicyRespBaseEffectMode struct {
	value string
}

type ApiPolicyRespBaseEffectModeEnum struct {
	ALL ApiPolicyRespBaseEffectMode
	ANY ApiPolicyRespBaseEffectMode
}

func GetApiPolicyRespBaseEffectModeEnum() ApiPolicyRespBaseEffectModeEnum {
	return ApiPolicyRespBaseEffectModeEnum{
		ALL: ApiPolicyRespBaseEffectMode{
			value: "ALL",
		},
		ANY: ApiPolicyRespBaseEffectMode{
			value: "ANY",
		},
	}
}

func (c ApiPolicyRespBaseEffectMode) Value() string {
	return c.value
}

func (c ApiPolicyRespBaseEffectMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyRespBaseEffectMode) UnmarshalJSON(b []byte) error {
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
