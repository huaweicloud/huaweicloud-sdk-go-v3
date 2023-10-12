package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApiRequest Request Object
type ShowApiRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType ShowApiRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// API ID
	ApiId string `json:"api_id"`
}

func (o ShowApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiRequest struct{}"
	}

	return strings.Join([]string{"ShowApiRequest", string(data)}, " ")
}

type ShowApiRequestDlmType struct {
	value string
}

type ShowApiRequestDlmTypeEnum struct {
	SHARED    ShowApiRequestDlmType
	EXCLUSIVE ShowApiRequestDlmType
}

func GetShowApiRequestDlmTypeEnum() ShowApiRequestDlmTypeEnum {
	return ShowApiRequestDlmTypeEnum{
		SHARED: ShowApiRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowApiRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowApiRequestDlmType) Value() string {
	return c.value
}

func (c ShowApiRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiRequestDlmType) UnmarshalJSON(b []byte) error {
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
