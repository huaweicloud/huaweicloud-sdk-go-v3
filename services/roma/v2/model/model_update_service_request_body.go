package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateServiceRequestBody struct {

	// 服务名称，支持中文,英文大小写，数字，下划线和中划线,长度2-64
	ServiceName string `json:"service_name" xml:"service_name"`

	// 服务描述，长度0-200
	Description *string `json:"description,omitempty" xml:"description"`

	// 服务状态 0-启用 1-停用
	Status UpdateServiceRequestBodyStatus `json:"status" xml:"status"`
}

func (o UpdateServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateServiceRequestBody", string(data)}, " ")
}

type UpdateServiceRequestBodyStatus struct {
	value int32
}

type UpdateServiceRequestBodyStatusEnum struct {
	E_0 UpdateServiceRequestBodyStatus
	E_1 UpdateServiceRequestBodyStatus
}

func GetUpdateServiceRequestBodyStatusEnum() UpdateServiceRequestBodyStatusEnum {
	return UpdateServiceRequestBodyStatusEnum{
		E_0: UpdateServiceRequestBodyStatus{
			value: 0,
		}, E_1: UpdateServiceRequestBodyStatus{
			value: 1,
		},
	}
}

func (c UpdateServiceRequestBodyStatus) Value() int32 {
	return c.value
}

func (c UpdateServiceRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateServiceRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
