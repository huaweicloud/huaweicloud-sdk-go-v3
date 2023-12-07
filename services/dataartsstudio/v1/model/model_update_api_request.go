package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateApiRequest Request Object
type UpdateApiRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *UpdateApiRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// API ID
	ApiId string `json:"api_id"`

	Body *Api `json:"body,omitempty"`
}

func (o UpdateApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApiRequest struct{}"
	}

	return strings.Join([]string{"UpdateApiRequest", string(data)}, " ")
}

type UpdateApiRequestDlmType struct {
	value string
}

type UpdateApiRequestDlmTypeEnum struct {
	SHARED    UpdateApiRequestDlmType
	EXCLUSIVE UpdateApiRequestDlmType
}

func GetUpdateApiRequestDlmTypeEnum() UpdateApiRequestDlmTypeEnum {
	return UpdateApiRequestDlmTypeEnum{
		SHARED: UpdateApiRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: UpdateApiRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c UpdateApiRequestDlmType) Value() string {
	return c.value
}

func (c UpdateApiRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateApiRequestDlmType) UnmarshalJSON(b []byte) error {
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
