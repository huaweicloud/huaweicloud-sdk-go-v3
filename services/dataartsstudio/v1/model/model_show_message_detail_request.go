package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMessageDetailRequest Request Object
type ShowMessageDetailRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *ShowMessageDetailRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 消息信息id
	MessageId string `json:"message_id"`
}

func (o ShowMessageDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessageDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMessageDetailRequest", string(data)}, " ")
}

type ShowMessageDetailRequestDlmType struct {
	value string
}

type ShowMessageDetailRequestDlmTypeEnum struct {
	SHARED    ShowMessageDetailRequestDlmType
	EXCLUSIVE ShowMessageDetailRequestDlmType
}

func GetShowMessageDetailRequestDlmTypeEnum() ShowMessageDetailRequestDlmTypeEnum {
	return ShowMessageDetailRequestDlmTypeEnum{
		SHARED: ShowMessageDetailRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowMessageDetailRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowMessageDetailRequestDlmType) Value() string {
	return c.value
}

func (c ShowMessageDetailRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMessageDetailRequestDlmType) UnmarshalJSON(b []byte) error {
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
