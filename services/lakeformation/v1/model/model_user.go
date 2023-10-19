package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// User user
type User struct {

	// 用户名
	UserName string `json:"user_name"`

	// 用户类型 IAM云用户 SAML联邦 LDAP ld用户 LOCAL 本地用户 AGENTTENANT 委托 OTHER 其它
	UserSource UserUserSource `json:"user_source"`

	// 用户ID
	UserId string `json:"user_id"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}

type UserUserSource struct {
	value string
}

type UserUserSourceEnum struct {
	IAM         UserUserSource
	SAML        UserUserSource
	LDAP        UserUserSource
	LOCAL       UserUserSource
	AGENTTENANT UserUserSource
	OTHER       UserUserSource
}

func GetUserUserSourceEnum() UserUserSourceEnum {
	return UserUserSourceEnum{
		IAM: UserUserSource{
			value: "IAM",
		},
		SAML: UserUserSource{
			value: "SAML",
		},
		LDAP: UserUserSource{
			value: "LDAP",
		},
		LOCAL: UserUserSource{
			value: "LOCAL",
		},
		AGENTTENANT: UserUserSource{
			value: "AGENTTENANT",
		},
		OTHER: UserUserSource{
			value: "OTHER",
		},
	}
}

func (c UserUserSource) Value() string {
	return c.value
}

func (c UserUserSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserUserSource) UnmarshalJSON(b []byte) error {
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
