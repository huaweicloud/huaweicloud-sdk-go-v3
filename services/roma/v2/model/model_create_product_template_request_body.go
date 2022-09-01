package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateProductTemplateRequestBody struct {

	// 产品模板名称，支持中文,英文大小写，数字，下划线和中划线,长度2-64
	Name string `json:"name" xml:"name"`

	// 产品模板描述，长度0-200
	Description *string `json:"description,omitempty" xml:"description"`

	// 产品模板状态 0-启用 1-禁用
	Status CreateProductTemplateRequestBodyStatus `json:"status" xml:"status"`
}

func (o CreateProductTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProductTemplateRequestBody", string(data)}, " ")
}

type CreateProductTemplateRequestBodyStatus struct {
	value int32
}

type CreateProductTemplateRequestBodyStatusEnum struct {
	E_0 CreateProductTemplateRequestBodyStatus
	E_1 CreateProductTemplateRequestBodyStatus
}

func GetCreateProductTemplateRequestBodyStatusEnum() CreateProductTemplateRequestBodyStatusEnum {
	return CreateProductTemplateRequestBodyStatusEnum{
		E_0: CreateProductTemplateRequestBodyStatus{
			value: 0,
		}, E_1: CreateProductTemplateRequestBodyStatus{
			value: 1,
		},
	}
}

func (c CreateProductTemplateRequestBodyStatus) Value() int32 {
	return c.value
}

func (c CreateProductTemplateRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateProductTemplateRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
