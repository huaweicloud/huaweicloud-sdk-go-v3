package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DebugApiRequest Request Object
type DebugApiRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType DebugApiRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// api编号
	ApiId string `json:"api_id"`

	// 集群编号
	InstanceId string `json:"instance_id"`

	Body *ApiTestDto `json:"body,omitempty"`
}

func (o DebugApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugApiRequest struct{}"
	}

	return strings.Join([]string{"DebugApiRequest", string(data)}, " ")
}

type DebugApiRequestDlmType struct {
	value string
}

type DebugApiRequestDlmTypeEnum struct {
	SHARED    DebugApiRequestDlmType
	EXCLUSIVE DebugApiRequestDlmType
}

func GetDebugApiRequestDlmTypeEnum() DebugApiRequestDlmTypeEnum {
	return DebugApiRequestDlmTypeEnum{
		SHARED: DebugApiRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: DebugApiRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c DebugApiRequestDlmType) Value() string {
	return c.value
}

func (c DebugApiRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DebugApiRequestDlmType) UnmarshalJSON(b []byte) error {
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
