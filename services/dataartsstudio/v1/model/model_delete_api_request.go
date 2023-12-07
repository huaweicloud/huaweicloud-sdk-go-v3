package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteApiRequest Request Object
type DeleteApiRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType *DeleteApiRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// API删除ID列表
	Body *[]string `json:"body,omitempty"`
}

func (o DeleteApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiRequest struct{}"
	}

	return strings.Join([]string{"DeleteApiRequest", string(data)}, " ")
}

type DeleteApiRequestDlmType struct {
	value string
}

type DeleteApiRequestDlmTypeEnum struct {
	SHARED    DeleteApiRequestDlmType
	EXCLUSIVE DeleteApiRequestDlmType
}

func GetDeleteApiRequestDlmTypeEnum() DeleteApiRequestDlmTypeEnum {
	return DeleteApiRequestDlmTypeEnum{
		SHARED: DeleteApiRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: DeleteApiRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c DeleteApiRequestDlmType) Value() string {
	return c.value
}

func (c DeleteApiRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteApiRequestDlmType) UnmarshalJSON(b []byte) error {
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
