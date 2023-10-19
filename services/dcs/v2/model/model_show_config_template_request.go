package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowConfigTemplateRequest Request Object
type ShowConfigTemplateRequest struct {

	// 模板ID
	TemplateId string `json:"template_id"`

	// 模板类型
	Type ShowConfigTemplateRequestType `json:"type"`
}

func (o ShowConfigTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigTemplateRequest", string(data)}, " ")
}

type ShowConfigTemplateRequestType struct {
	value string
}

type ShowConfigTemplateRequestTypeEnum struct {
	SYS  ShowConfigTemplateRequestType
	USER ShowConfigTemplateRequestType
}

func GetShowConfigTemplateRequestTypeEnum() ShowConfigTemplateRequestTypeEnum {
	return ShowConfigTemplateRequestTypeEnum{
		SYS: ShowConfigTemplateRequestType{
			value: "sys",
		},
		USER: ShowConfigTemplateRequestType{
			value: "user",
		},
	}
}

func (c ShowConfigTemplateRequestType) Value() string {
	return c.value
}

func (c ShowConfigTemplateRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConfigTemplateRequestType) UnmarshalJSON(b []byte) error {
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
