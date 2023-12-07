package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteServiceCatalogRequest Request Object
type DeleteServiceCatalogRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *DeleteServiceCatalogRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	Body *ApiCatalogDeleteParaDto `json:"body,omitempty"`
}

func (o DeleteServiceCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceCatalogRequest struct{}"
	}

	return strings.Join([]string{"DeleteServiceCatalogRequest", string(data)}, " ")
}

type DeleteServiceCatalogRequestDlmType struct {
	value string
}

type DeleteServiceCatalogRequestDlmTypeEnum struct {
	SHARED    DeleteServiceCatalogRequestDlmType
	EXCLUSIVE DeleteServiceCatalogRequestDlmType
}

func GetDeleteServiceCatalogRequestDlmTypeEnum() DeleteServiceCatalogRequestDlmTypeEnum {
	return DeleteServiceCatalogRequestDlmTypeEnum{
		SHARED: DeleteServiceCatalogRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: DeleteServiceCatalogRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c DeleteServiceCatalogRequestDlmType) Value() string {
	return c.value
}

func (c DeleteServiceCatalogRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteServiceCatalogRequestDlmType) UnmarshalJSON(b []byte) error {
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
