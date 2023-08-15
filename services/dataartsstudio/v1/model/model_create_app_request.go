package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAppRequest Request Object
type CreateAppRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType CreateAppRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	Body *AppRequestDto `json:"body,omitempty"`
}

func (o CreateAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppRequest struct{}"
	}

	return strings.Join([]string{"CreateAppRequest", string(data)}, " ")
}

type CreateAppRequestDlmType struct {
	value string
}

type CreateAppRequestDlmTypeEnum struct {
	SHARED    CreateAppRequestDlmType
	EXCLUSIVE CreateAppRequestDlmType
}

func GetCreateAppRequestDlmTypeEnum() CreateAppRequestDlmTypeEnum {
	return CreateAppRequestDlmTypeEnum{
		SHARED: CreateAppRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: CreateAppRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c CreateAppRequestDlmType) Value() string {
	return c.value
}

func (c CreateAppRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAppRequestDlmType) UnmarshalJSON(b []byte) error {
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
