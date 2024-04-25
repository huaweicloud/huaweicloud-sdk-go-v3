package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateServiceCatalogRequest Request Object
type CreateServiceCatalogRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *CreateServiceCatalogRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	Body *ApiCatalogCreateParaDto `json:"body,omitempty"`
}

func (o CreateServiceCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceCatalogRequest struct{}"
	}

	return strings.Join([]string{"CreateServiceCatalogRequest", string(data)}, " ")
}

type CreateServiceCatalogRequestDlmType struct {
	value string
}

type CreateServiceCatalogRequestDlmTypeEnum struct {
	SHARED    CreateServiceCatalogRequestDlmType
	EXCLUSIVE CreateServiceCatalogRequestDlmType
}

func GetCreateServiceCatalogRequestDlmTypeEnum() CreateServiceCatalogRequestDlmTypeEnum {
	return CreateServiceCatalogRequestDlmTypeEnum{
		SHARED: CreateServiceCatalogRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: CreateServiceCatalogRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c CreateServiceCatalogRequestDlmType) Value() string {
	return c.value
}

func (c CreateServiceCatalogRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateServiceCatalogRequestDlmType) UnmarshalJSON(b []byte) error {
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
