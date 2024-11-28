package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateCbcOrderRequestBody struct {

	// change_mode
	ChangeMode int32 `json:"change_mode"`

	// 该请求参数为预留参数，当前不支持自动支付，使用接口时该参数请使用0 0：不自动支付 1：自动支付
	IsAutoPay UpdateCbcOrderRequestBodyIsAutoPay `json:"is_auto_pay"`

	// 发起规格变更操作的云服务类型
	CloudServiceType string `json:"cloud_service_type"`

	// project_id
	ProjectId string `json:"project_id"`

	// 资源标识ID
	ResourceId string `json:"resource_id"`

	ProductInfo *UpdateCbcOrderRequestBodyProductInfo `json:"product_info"`
}

func (o UpdateCbcOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCbcOrderRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCbcOrderRequestBody", string(data)}, " ")
}

type UpdateCbcOrderRequestBodyIsAutoPay struct {
	value int32
}

type UpdateCbcOrderRequestBodyIsAutoPayEnum struct {
	E_0 UpdateCbcOrderRequestBodyIsAutoPay
	E_1 UpdateCbcOrderRequestBodyIsAutoPay
}

func GetUpdateCbcOrderRequestBodyIsAutoPayEnum() UpdateCbcOrderRequestBodyIsAutoPayEnum {
	return UpdateCbcOrderRequestBodyIsAutoPayEnum{
		E_0: UpdateCbcOrderRequestBodyIsAutoPay{
			value: 0,
		}, E_1: UpdateCbcOrderRequestBodyIsAutoPay{
			value: 1,
		},
	}
}

func (c UpdateCbcOrderRequestBodyIsAutoPay) Value() int32 {
	return c.value
}

func (c UpdateCbcOrderRequestBodyIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCbcOrderRequestBodyIsAutoPay) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
