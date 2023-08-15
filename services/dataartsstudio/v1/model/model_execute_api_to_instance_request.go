package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteApiToInstanceRequest Request Object
type ExecuteApiToInstanceRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType ExecuteApiToInstanceRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// api编号
	ApiId string `json:"api_id"`

	// 集群编号
	InstanceId string `json:"instance_id"`

	Body *ApiActionDto `json:"body,omitempty"`
}

func (o ExecuteApiToInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteApiToInstanceRequest struct{}"
	}

	return strings.Join([]string{"ExecuteApiToInstanceRequest", string(data)}, " ")
}

type ExecuteApiToInstanceRequestDlmType struct {
	value string
}

type ExecuteApiToInstanceRequestDlmTypeEnum struct {
	SHARED    ExecuteApiToInstanceRequestDlmType
	EXCLUSIVE ExecuteApiToInstanceRequestDlmType
}

func GetExecuteApiToInstanceRequestDlmTypeEnum() ExecuteApiToInstanceRequestDlmTypeEnum {
	return ExecuteApiToInstanceRequestDlmTypeEnum{
		SHARED: ExecuteApiToInstanceRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ExecuteApiToInstanceRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ExecuteApiToInstanceRequestDlmType) Value() string {
	return c.value
}

func (c ExecuteApiToInstanceRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteApiToInstanceRequestDlmType) UnmarshalJSON(b []byte) error {
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
