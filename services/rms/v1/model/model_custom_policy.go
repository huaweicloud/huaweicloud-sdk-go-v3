package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CustomPolicy 自定义合规规则
type CustomPolicy struct {

	// 自定义函数的urn
	FunctionUrn string `json:"function_urn"`

	// 自定义合规规则调用function方式
	AuthType CustomPolicyAuthType `json:"auth_type"`

	// method参数值，method为agency时，为{\"agency_name\":rms_fg_agency}, rms_fg_agency为授权RMS调用FunctionGraph接口的委托名称
	AuthValue map[string]interface{} `json:"auth_value,omitempty"`
}

func (o CustomPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomPolicy struct{}"
	}

	return strings.Join([]string{"CustomPolicy", string(data)}, " ")
}

type CustomPolicyAuthType struct {
	value string
}

type CustomPolicyAuthTypeEnum struct {
	AGENCY CustomPolicyAuthType
}

func GetCustomPolicyAuthTypeEnum() CustomPolicyAuthTypeEnum {
	return CustomPolicyAuthTypeEnum{
		AGENCY: CustomPolicyAuthType{
			value: "agency",
		},
	}
}

func (c CustomPolicyAuthType) Value() string {
	return c.value
}

func (c CustomPolicyAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomPolicyAuthType) UnmarshalJSON(b []byte) error {
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
