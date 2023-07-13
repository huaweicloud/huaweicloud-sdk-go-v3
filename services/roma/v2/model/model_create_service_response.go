package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateServiceResponse Response Object
type CreateServiceResponse struct {

	// 服务归属的产品模板ID
	ProductTemplateId *int32 `json:"product_template_id,omitempty"`

	// 服务归属的产品ID
	ProductId *int32 `json:"product_id,omitempty"`

	// 服务ID
	ServiceId *int32 `json:"service_id,omitempty"`

	// 服务名称，支持中文,英文大小写，数字，下划线和中划线,长度2-64
	ServiceName *string `json:"service_name,omitempty"`

	// 服务描述，长度0-200
	Description *string `json:"description,omitempty"`

	// 服务状态 0-启用 1-停用
	Status *CreateServiceResponseStatus `json:"status,omitempty"`

	CreatedUser *CreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty"`

	// 创建时间止，格式timestamp(ms)，使用UTC时区
	CreatedDatetime *int64 `json:"created_datetime,omitempty"`

	// 最后修改时间止，格式timestamp(ms)，使用UTC时区
	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o CreateServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceResponse struct{}"
	}

	return strings.Join([]string{"CreateServiceResponse", string(data)}, " ")
}

type CreateServiceResponseStatus struct {
	value int32
}

type CreateServiceResponseStatusEnum struct {
	E_0 CreateServiceResponseStatus
	E_1 CreateServiceResponseStatus
}

func GetCreateServiceResponseStatusEnum() CreateServiceResponseStatusEnum {
	return CreateServiceResponseStatusEnum{
		E_0: CreateServiceResponseStatus{
			value: 0,
		}, E_1: CreateServiceResponseStatus{
			value: 1,
		},
	}
}

func (c CreateServiceResponseStatus) Value() int32 {
	return c.value
}

func (c CreateServiceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateServiceResponseStatus) UnmarshalJSON(b []byte) error {
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
