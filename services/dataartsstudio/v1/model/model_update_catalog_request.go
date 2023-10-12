package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateCatalogRequest Request Object
type UpdateCatalogRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType UpdateCatalogRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 目录编号
	CatalogId string `json:"catalog_id"`

	Body *ApiCatalogUpdateParaDto `json:"body,omitempty"`
}

func (o UpdateCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCatalogRequest struct{}"
	}

	return strings.Join([]string{"UpdateCatalogRequest", string(data)}, " ")
}

type UpdateCatalogRequestDlmType struct {
	value string
}

type UpdateCatalogRequestDlmTypeEnum struct {
	SHARED    UpdateCatalogRequestDlmType
	EXCLUSIVE UpdateCatalogRequestDlmType
}

func GetUpdateCatalogRequestDlmTypeEnum() UpdateCatalogRequestDlmTypeEnum {
	return UpdateCatalogRequestDlmTypeEnum{
		SHARED: UpdateCatalogRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: UpdateCatalogRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c UpdateCatalogRequestDlmType) Value() string {
	return c.value
}

func (c UpdateCatalogRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCatalogRequestDlmType) UnmarshalJSON(b []byte) error {
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
