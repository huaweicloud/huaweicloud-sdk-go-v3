package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAppRequest Request Object
type UpdateAppRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *UpdateAppRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 应用编号
	AppId string `json:"app_id"`

	Body *AppUpdateDto `json:"body,omitempty"`
}

func (o UpdateAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppRequest", string(data)}, " ")
}

type UpdateAppRequestDlmType struct {
	value string
}

type UpdateAppRequestDlmTypeEnum struct {
	SHARED    UpdateAppRequestDlmType
	EXCLUSIVE UpdateAppRequestDlmType
}

func GetUpdateAppRequestDlmTypeEnum() UpdateAppRequestDlmTypeEnum {
	return UpdateAppRequestDlmTypeEnum{
		SHARED: UpdateAppRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: UpdateAppRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c UpdateAppRequestDlmType) Value() string {
	return c.value
}

func (c UpdateAppRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAppRequestDlmType) UnmarshalJSON(b []byte) error {
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
