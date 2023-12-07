package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowCatalogDetailRequest Request Object
type ShowCatalogDetailRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *ShowCatalogDetailRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 目录编号
	CatalogId string `json:"catalog_id"`
}

func (o ShowCatalogDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCatalogDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowCatalogDetailRequest", string(data)}, " ")
}

type ShowCatalogDetailRequestDlmType struct {
	value string
}

type ShowCatalogDetailRequestDlmTypeEnum struct {
	SHARED    ShowCatalogDetailRequestDlmType
	EXCLUSIVE ShowCatalogDetailRequestDlmType
}

func GetShowCatalogDetailRequestDlmTypeEnum() ShowCatalogDetailRequestDlmTypeEnum {
	return ShowCatalogDetailRequestDlmTypeEnum{
		SHARED: ShowCatalogDetailRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowCatalogDetailRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowCatalogDetailRequestDlmType) Value() string {
	return c.value
}

func (c ShowCatalogDetailRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCatalogDetailRequestDlmType) UnmarshalJSON(b []byte) error {
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
