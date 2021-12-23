package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPolicyMockResp struct {
	// 编号

	Id *string `json:"id,omitempty"`
	// 关联的策略组合模式： - ALL：满足全部条件 - ANY：满足任一条件

	EffectMode ApiPolicyMockRespEffectMode `json:"effect_mode"`
	// 策略后端名称。字符串由中文、英文字母、数字、下划线组成，且只能以中文或英文开头。

	Name string `json:"name"`
	// 后端参数列表

	BackendParams *[]BackendParam `json:"backend_params,omitempty"`
	// 策略条件列表

	Conditions []CoditionResp `json:"conditions"`
	// 后端自定义认证对象的ID

	AuthorizerId *string `json:"authorizer_id,omitempty"`
	// 返回结果

	ResultContent *string `json:"result_content,omitempty"`
}

func (o ApiPolicyMockResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyMockResp struct{}"
	}

	return strings.Join([]string{"ApiPolicyMockResp", string(data)}, " ")
}

type ApiPolicyMockRespEffectMode struct {
	value string
}

type ApiPolicyMockRespEffectModeEnum struct {
	ALL ApiPolicyMockRespEffectMode
	ANY ApiPolicyMockRespEffectMode
}

func GetApiPolicyMockRespEffectModeEnum() ApiPolicyMockRespEffectModeEnum {
	return ApiPolicyMockRespEffectModeEnum{
		ALL: ApiPolicyMockRespEffectMode{
			value: "ALL",
		},
		ANY: ApiPolicyMockRespEffectMode{
			value: "ANY",
		},
	}
}

func (c ApiPolicyMockRespEffectMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyMockRespEffectMode) UnmarshalJSON(b []byte) error {
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
