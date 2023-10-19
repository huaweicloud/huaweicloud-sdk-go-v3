package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInstanceRequestBody 创建LakeFormation实例请求体
type CreateInstanceRequestBody struct {

	// 实例名称。只能包含字母、数字、下划线和中划线，且长度为4到32个字符。
	Name string `json:"name"`

	// 支付类型，postPaid为按需期，prePaid为包周期
	ChargeMode CreateInstanceRequestBodyChargeMode `json:"charge_mode"`

	// 企业项目ID，只有对接了企业项目才可以填写。只能包含字母、数字和中划线，且长度为1到64个字符。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例描述。用户输入的描述，最长为255个字符。
	Description *string `json:"description,omitempty"`

	// false为物理多租；true为逻辑多租。默认为true。
	Shared bool `json:"shared"`

	// 包周期订购时的订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 规格列表
	Specs *[]CreateSpec `json:"specs,omitempty"`

	// 标签列表，最多添加20个标签。
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
	PRE_PAID  CreateInstanceRequestBodyChargeMode
}

func GetCreateInstanceRequestBodyChargeModeEnum() CreateInstanceRequestBodyChargeModeEnum {
	return CreateInstanceRequestBodyChargeModeEnum{
		POST_PAID: CreateInstanceRequestBodyChargeMode{
			value: "postPaid",
		},
		PRE_PAID: CreateInstanceRequestBodyChargeMode{
			value: "prePaid",
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
