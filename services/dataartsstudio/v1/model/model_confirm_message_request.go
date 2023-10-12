package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfirmMessageRequest Request Object
type ConfirmMessageRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType ConfirmMessageRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	Body *OpenApiParaForCheckMessage `json:"body,omitempty"`
}

func (o ConfirmMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmMessageRequest struct{}"
	}

	return strings.Join([]string{"ConfirmMessageRequest", string(data)}, " ")
}

type ConfirmMessageRequestDlmType struct {
	value string
}

type ConfirmMessageRequestDlmTypeEnum struct {
	SHARED    ConfirmMessageRequestDlmType
	EXCLUSIVE ConfirmMessageRequestDlmType
}

func GetConfirmMessageRequestDlmTypeEnum() ConfirmMessageRequestDlmTypeEnum {
	return ConfirmMessageRequestDlmTypeEnum{
		SHARED: ConfirmMessageRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ConfirmMessageRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ConfirmMessageRequestDlmType) Value() string {
	return c.value
}

func (c ConfirmMessageRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmMessageRequestDlmType) UnmarshalJSON(b []byte) error {
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
