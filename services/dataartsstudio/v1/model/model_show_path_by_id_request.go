package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPathByIdRequest Request Object
type ShowPathByIdRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *ShowPathByIdRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 目录编号
	CatalogId string `json:"catalog_id"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowPathByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPathByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowPathByIdRequest", string(data)}, " ")
}

type ShowPathByIdRequestDlmType struct {
	value string
}

type ShowPathByIdRequestDlmTypeEnum struct {
	SHARED    ShowPathByIdRequestDlmType
	EXCLUSIVE ShowPathByIdRequestDlmType
}

func GetShowPathByIdRequestDlmTypeEnum() ShowPathByIdRequestDlmTypeEnum {
	return ShowPathByIdRequestDlmTypeEnum{
		SHARED: ShowPathByIdRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowPathByIdRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowPathByIdRequestDlmType) Value() string {
	return c.value
}

func (c ShowPathByIdRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPathByIdRequestDlmType) UnmarshalJSON(b []byte) error {
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
