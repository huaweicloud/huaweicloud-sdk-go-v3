package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type CreateServiceRequestBody struct {
	// 服务归属的产品模板ID，产品模板ID和产品ID二选一必填，自动向下取整

	ProductTemplateId *int32 `json:"product_template_id,omitempty"`
	// 服务归属的产品ID，产品模板ID和产品ID二选一必填，自动向下取整

	ProductId *int32 `json:"product_id,omitempty"`
	// 服务名称，支持中文,英文大小写，数字，下划线和中划线,长度2-64

	ServiceName string `json:"service_name"`
	// 服务描述，长度0-200

	Description *string `json:"description,omitempty"`
	// 服务状态 0-启用 1-停用

	Status CreateServiceRequestBodyStatus `json:"status"`
}

func (o CreateServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateServiceRequestBody", string(data)}, " ")
}

type CreateServiceRequestBodyStatus struct {
	value int32
}

type CreateServiceRequestBodyStatusEnum struct {
	E_0 CreateServiceRequestBodyStatus
	E_1 CreateServiceRequestBodyStatus
}

func GetCreateServiceRequestBodyStatusEnum() CreateServiceRequestBodyStatusEnum {
	return CreateServiceRequestBodyStatusEnum{
		E_0: CreateServiceRequestBodyStatus{
			value: 0,
		}, E_1: CreateServiceRequestBodyStatus{
			value: 1,
		},
	}
}

func (c CreateServiceRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateServiceRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
