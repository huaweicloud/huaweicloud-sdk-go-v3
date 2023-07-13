package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Service struct {

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
	Status *ServiceStatus `json:"status,omitempty"`

	CreatedUser *CreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty"`

	// 创建时间止，格式timestamp(ms)，使用UTC时区
	CreatedDatetime *int64 `json:"created_datetime,omitempty"`

	// 最后修改时间止，格式timestamp(ms)，使用UTC时区
	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty"`
}

func (o Service) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Service struct{}"
	}

	return strings.Join([]string{"Service", string(data)}, " ")
}

type ServiceStatus struct {
	value int32
}

type ServiceStatusEnum struct {
	E_0 ServiceStatus
	E_1 ServiceStatus
}

func GetServiceStatusEnum() ServiceStatusEnum {
	return ServiceStatusEnum{
		E_0: ServiceStatus{
			value: 0,
		}, E_1: ServiceStatus{
			value: 1,
		},
	}
}

func (c ServiceStatus) Value() int32 {
	return c.value
}

func (c ServiceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServiceStatus) UnmarshalJSON(b []byte) error {
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
