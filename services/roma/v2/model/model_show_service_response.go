package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowServiceResponse struct {
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

	Status *ShowServiceResponseStatus `json:"status,omitempty"`

	CreatedUser *CreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty"`
	// 创建时间止，格式timestamp(ms)，使用UTC时区

	CreatedDatetime *int64 `json:"created_datetime,omitempty"`
	// 最后修改时间止，格式timestamp(ms)，使用UTC时区

	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o ShowServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServiceResponse struct{}"
	}

	return strings.Join([]string{"ShowServiceResponse", string(data)}, " ")
}

type ShowServiceResponseStatus struct {
	value int32
}

type ShowServiceResponseStatusEnum struct {
	E_0 ShowServiceResponseStatus
	E_1 ShowServiceResponseStatus
}

func GetShowServiceResponseStatusEnum() ShowServiceResponseStatusEnum {
	return ShowServiceResponseStatusEnum{
		E_0: ShowServiceResponseStatus{
			value: 0,
		}, E_1: ShowServiceResponseStatus{
			value: 1,
		},
	}
}

func (c ShowServiceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowServiceResponseStatus) UnmarshalJSON(b []byte) error {
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
