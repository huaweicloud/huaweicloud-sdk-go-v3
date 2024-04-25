package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteApiToInstanceRequest Request Object
type ExecuteApiToInstanceRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ExecuteApiToInstanceRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// api编号。
	ApiId string `json:"api_id"`

	// 集群编号。
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
