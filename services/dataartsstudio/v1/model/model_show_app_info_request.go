package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAppInfoRequest Request Object
type ShowAppInfoRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType ShowAppInfoRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 应用编号
	AppId string `json:"app_id"`
}

func (o ShowAppInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowAppInfoRequest", string(data)}, " ")
}

type ShowAppInfoRequestDlmType struct {
	value string
}

type ShowAppInfoRequestDlmTypeEnum struct {
	SHARED    ShowAppInfoRequestDlmType
	EXCLUSIVE ShowAppInfoRequestDlmType
}

func GetShowAppInfoRequestDlmTypeEnum() ShowAppInfoRequestDlmTypeEnum {
	return ShowAppInfoRequestDlmTypeEnum{
		SHARED: ShowAppInfoRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowAppInfoRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowAppInfoRequestDlmType) Value() string {
	return c.value
}

func (c ShowAppInfoRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppInfoRequestDlmType) UnmarshalJSON(b []byte) error {
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
