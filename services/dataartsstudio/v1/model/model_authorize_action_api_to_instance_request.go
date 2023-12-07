package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AuthorizeActionApiToInstanceRequest Request Object
type AuthorizeActionApiToInstanceRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *AuthorizeActionApiToInstanceRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	Body *ApiParaForAuthToInstance `json:"body,omitempty"`
}

func (o AuthorizeActionApiToInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeActionApiToInstanceRequest struct{}"
	}

	return strings.Join([]string{"AuthorizeActionApiToInstanceRequest", string(data)}, " ")
}

type AuthorizeActionApiToInstanceRequestDlmType struct {
	value string
}

type AuthorizeActionApiToInstanceRequestDlmTypeEnum struct {
	SHARED    AuthorizeActionApiToInstanceRequestDlmType
	EXCLUSIVE AuthorizeActionApiToInstanceRequestDlmType
}

func GetAuthorizeActionApiToInstanceRequestDlmTypeEnum() AuthorizeActionApiToInstanceRequestDlmTypeEnum {
	return AuthorizeActionApiToInstanceRequestDlmTypeEnum{
		SHARED: AuthorizeActionApiToInstanceRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: AuthorizeActionApiToInstanceRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c AuthorizeActionApiToInstanceRequestDlmType) Value() string {
	return c.value
}

func (c AuthorizeActionApiToInstanceRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizeActionApiToInstanceRequestDlmType) UnmarshalJSON(b []byte) error {
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
