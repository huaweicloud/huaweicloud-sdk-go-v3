package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AuthorizeApiToInstanceRequest Request Object
type AuthorizeApiToInstanceRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *AuthorizeApiToInstanceRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// api编号。
	ApiId string `json:"api_id"`

	// 集群编号。
	InstanceId string `json:"instance_id"`

	Body *ApiParaForAuthorizeToInstance `json:"body,omitempty"`
}

func (o AuthorizeApiToInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeApiToInstanceRequest struct{}"
	}

	return strings.Join([]string{"AuthorizeApiToInstanceRequest", string(data)}, " ")
}

type AuthorizeApiToInstanceRequestDlmType struct {
	value string
}

type AuthorizeApiToInstanceRequestDlmTypeEnum struct {
	SHARED    AuthorizeApiToInstanceRequestDlmType
	EXCLUSIVE AuthorizeApiToInstanceRequestDlmType
}

func GetAuthorizeApiToInstanceRequestDlmTypeEnum() AuthorizeApiToInstanceRequestDlmTypeEnum {
	return AuthorizeApiToInstanceRequestDlmTypeEnum{
		SHARED: AuthorizeApiToInstanceRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: AuthorizeApiToInstanceRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c AuthorizeApiToInstanceRequestDlmType) Value() string {
	return c.value
}

func (c AuthorizeApiToInstanceRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizeApiToInstanceRequestDlmType) UnmarshalJSON(b []byte) error {
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
