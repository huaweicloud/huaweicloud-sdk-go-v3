package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DebugApiRequest Request Object
type DebugApiRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *DebugApiRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// api编号。
	ApiId string `json:"api_id"`

	// 集群编号。
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
