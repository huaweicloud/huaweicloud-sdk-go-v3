package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PublishApiToInstanceRequest Request Object
type PublishApiToInstanceRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType PublishApiToInstanceRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// api编号
	ApiId string `json:"api_id"`

	// 集群编号
	InstanceId string `json:"instance_id"`

	Body *GatewayDto `json:"body,omitempty"`
}

func (o PublishApiToInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishApiToInstanceRequest struct{}"
	}

	return strings.Join([]string{"PublishApiToInstanceRequest", string(data)}, " ")
}

type PublishApiToInstanceRequestDlmType struct {
	value string
}

type PublishApiToInstanceRequestDlmTypeEnum struct {
	SHARED    PublishApiToInstanceRequestDlmType
	EXCLUSIVE PublishApiToInstanceRequestDlmType
}

func GetPublishApiToInstanceRequestDlmTypeEnum() PublishApiToInstanceRequestDlmTypeEnum {
	return PublishApiToInstanceRequestDlmTypeEnum{
		SHARED: PublishApiToInstanceRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: PublishApiToInstanceRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c PublishApiToInstanceRequestDlmType) Value() string {
	return c.value
}

func (c PublishApiToInstanceRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublishApiToInstanceRequestDlmType) UnmarshalJSON(b []byte) error {
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
