package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PublishApiRequest Request Object
type PublishApiRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType PublishApiRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	Body *OpenApiParaForPublish `json:"body,omitempty"`
}

func (o PublishApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishApiRequest struct{}"
	}

	return strings.Join([]string{"PublishApiRequest", string(data)}, " ")
}

type PublishApiRequestDlmType struct {
	value string
}

type PublishApiRequestDlmTypeEnum struct {
	SHARED    PublishApiRequestDlmType
	EXCLUSIVE PublishApiRequestDlmType
}

func GetPublishApiRequestDlmTypeEnum() PublishApiRequestDlmTypeEnum {
	return PublishApiRequestDlmTypeEnum{
		SHARED: PublishApiRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: PublishApiRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c PublishApiRequestDlmType) Value() string {
	return c.value
}

func (c PublishApiRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublishApiRequestDlmType) UnmarshalJSON(b []byte) error {
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
