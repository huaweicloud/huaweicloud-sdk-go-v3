package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateServiceCatalogRequest Request Object
type CreateServiceCatalogRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *CreateServiceCatalogRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
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
