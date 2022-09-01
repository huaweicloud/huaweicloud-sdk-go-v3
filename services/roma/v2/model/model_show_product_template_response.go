package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowProductTemplateResponse struct {

	// 产品模板ID
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 产品模板名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 产品模板描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 产品模板状态 0-启用 1-停用
	Status *ShowProductTemplateResponseStatus `json:"status,omitempty" xml:"status"`

	CreatedUser *CreatedUser `json:"created_user,omitempty" xml:"created_user"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty" xml:"last_updated_user"`

	// 创建时间，timestamp(ms)，使用UTC时区
	CreatedDatetime *int64 `json:"created_datetime,omitempty" xml:"created_datetime"`

	// 最后修改时间，timestamp(ms)，使用UTC时区
	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty" xml:"last_updated_datetime"`
	HttpStatusCode      int    `json:"-"`
}

func (o ShowProductTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProductTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowProductTemplateResponse", string(data)}, " ")
}

type ShowProductTemplateResponseStatus struct {
	value int32
}

type ShowProductTemplateResponseStatusEnum struct {
	E_0 ShowProductTemplateResponseStatus
	E_1 ShowProductTemplateResponseStatus
}

func GetShowProductTemplateResponseStatusEnum() ShowProductTemplateResponseStatusEnum {
	return ShowProductTemplateResponseStatusEnum{
		E_0: ShowProductTemplateResponseStatus{
			value: 0,
		}, E_1: ShowProductTemplateResponseStatus{
			value: 1,
		},
	}
}

func (c ShowProductTemplateResponseStatus) Value() int32 {
	return c.value
}

func (c ShowProductTemplateResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProductTemplateResponseStatus) UnmarshalJSON(b []byte) error {
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
