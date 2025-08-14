package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SignInOptionsDto 门户登录选项
type SignInOptionsDto struct {

	// IAM身份中心跳转应用程序的方式
	Origin SignInOptionsDtoOrigin `json:"origin"`

	// 接受应用程序身份验证请求的Url
	ApplicationUrl *string `json:"application_url,omitempty"`
}

func (o SignInOptionsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignInOptionsDto struct{}"
	}

	return strings.Join([]string{"SignInOptionsDto", string(data)}, " ")
}

type SignInOptionsDtoOrigin struct {
	value string
}

type SignInOptionsDtoOriginEnum struct {
	IDENTITY_CENTER SignInOptionsDtoOrigin
	APPLICATION     SignInOptionsDtoOrigin
}

func GetSignInOptionsDtoOriginEnum() SignInOptionsDtoOriginEnum {
	return SignInOptionsDtoOriginEnum{
		IDENTITY_CENTER: SignInOptionsDtoOrigin{
			value: "IDENTITY_CENTER",
		},
		APPLICATION: SignInOptionsDtoOrigin{
			value: "APPLICATION",
		},
	}
}

func (c SignInOptionsDtoOrigin) Value() string {
	return c.value
}

func (c SignInOptionsDtoOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignInOptionsDtoOrigin) UnmarshalJSON(b []byte) error {
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
