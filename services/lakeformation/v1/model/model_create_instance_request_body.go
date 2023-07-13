package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInstanceRequestBody 创建LakeFormation实例请求体
type CreateInstanceRequestBody struct {

	// 实例名
	Name string `json:"name"`

	// 支付类型，postPaid为按需期
	ChargeMode CreateInstanceRequestBodyChargeMode `json:"charge_mode"`

	// 企业项目id，只有对接了企业项目才可以填写
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// false:物理多租 true:逻辑多租
	Shared bool `json:"shared"`

	// 规格列表
	Specs []CreateSpec `json:"specs"`

	// 标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o CreateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequestBody", string(data)}, " ")
}

type CreateInstanceRequestBodyChargeMode struct {
	value string
}

type CreateInstanceRequestBodyChargeModeEnum struct {
	POST_PAID CreateInstanceRequestBodyChargeMode
}

func GetCreateInstanceRequestBodyChargeModeEnum() CreateInstanceRequestBodyChargeModeEnum {
	return CreateInstanceRequestBodyChargeModeEnum{
		POST_PAID: CreateInstanceRequestBodyChargeMode{
			value: "postPaid",
		},
	}
}

func (c CreateInstanceRequestBodyChargeMode) Value() string {
	return c.value
}

func (c CreateInstanceRequestBodyChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceRequestBodyChargeMode) UnmarshalJSON(b []byte) error {
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
