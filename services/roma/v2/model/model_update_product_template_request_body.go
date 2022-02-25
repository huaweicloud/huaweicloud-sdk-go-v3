package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateProductTemplateRequestBody struct {
	// 产品模板名称，支持中文,英文大小写，数字，下划线和中划线,长度2-64

	Name string `json:"name"`
	// 产品模板描述，长度0-200

	Description *string `json:"description,omitempty"`
	// 产品模板状态 0-启用 1-禁用

	Status UpdateProductTemplateRequestBodyStatus `json:"status"`
}

func (o UpdateProductTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProductTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProductTemplateRequestBody", string(data)}, " ")
}

type UpdateProductTemplateRequestBodyStatus struct {
	value int32
}

type UpdateProductTemplateRequestBodyStatusEnum struct {
	E_0 UpdateProductTemplateRequestBodyStatus
	E_1 UpdateProductTemplateRequestBodyStatus
}

func GetUpdateProductTemplateRequestBodyStatusEnum() UpdateProductTemplateRequestBodyStatusEnum {
	return UpdateProductTemplateRequestBodyStatusEnum{
		E_0: UpdateProductTemplateRequestBodyStatus{
			value: 0,
		}, E_1: UpdateProductTemplateRequestBodyStatus{
			value: 1,
		},
	}
}

func (c UpdateProductTemplateRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateProductTemplateRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
