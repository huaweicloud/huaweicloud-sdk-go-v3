package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MigrateApiRequest Request Object
type MigrateApiRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType MigrateApiRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	Body *ApiMoveParaDto `json:"body,omitempty"`
}

func (o MigrateApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateApiRequest struct{}"
	}

	return strings.Join([]string{"MigrateApiRequest", string(data)}, " ")
}

type MigrateApiRequestDlmType struct {
	value string
}

type MigrateApiRequestDlmTypeEnum struct {
	SHARED    MigrateApiRequestDlmType
	EXCLUSIVE MigrateApiRequestDlmType
}

func GetMigrateApiRequestDlmTypeEnum() MigrateApiRequestDlmTypeEnum {
	return MigrateApiRequestDlmTypeEnum{
		SHARED: MigrateApiRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: MigrateApiRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c MigrateApiRequestDlmType) Value() string {
	return c.value
}

func (c MigrateApiRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrateApiRequestDlmType) UnmarshalJSON(b []byte) error {
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
