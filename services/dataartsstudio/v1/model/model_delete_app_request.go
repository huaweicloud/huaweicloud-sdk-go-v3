package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteAppRequest Request Object
type DeleteAppRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType DeleteAppRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 应用编号
	AppId string `json:"app_id"`
}

func (o DeleteAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppRequest", string(data)}, " ")
}

type DeleteAppRequestDlmType struct {
	value string
}

type DeleteAppRequestDlmTypeEnum struct {
	SHARED    DeleteAppRequestDlmType
	EXCLUSIVE DeleteAppRequestDlmType
}

func GetDeleteAppRequestDlmTypeEnum() DeleteAppRequestDlmTypeEnum {
	return DeleteAppRequestDlmTypeEnum{
		SHARED: DeleteAppRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: DeleteAppRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c DeleteAppRequestDlmType) Value() string {
	return c.value
}

func (c DeleteAppRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteAppRequestDlmType) UnmarshalJSON(b []byte) error {
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
