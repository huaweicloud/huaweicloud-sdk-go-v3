package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UserGroup usergroup
type UserGroup struct {

	// 用户组名
	GroupName string `json:"group_name"`

	// 用户组类型
	GroupSource UserGroupGroupSource `json:"group_source"`

	// 用户组id
	GroupId string `json:"group_id"`
}

func (o UserGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserGroup struct{}"
	}

	return strings.Join([]string{"UserGroup", string(data)}, " ")
}

type UserGroupGroupSource struct {
	value string
}

type UserGroupGroupSourceEnum struct {
	IAM   UserGroupGroupSource
	SAML  UserGroupGroupSource
	LDAP  UserGroupGroupSource
	LOCAL UserGroupGroupSource
	OTHER UserGroupGroupSource
}

func GetUserGroupGroupSourceEnum() UserGroupGroupSourceEnum {
	return UserGroupGroupSourceEnum{
		IAM: UserGroupGroupSource{
			value: "IAM",
		},
		SAML: UserGroupGroupSource{
			value: "SAML",
		},
		LDAP: UserGroupGroupSource{
			value: "LDAP",
		},
		LOCAL: UserGroupGroupSource{
			value: "LOCAL",
		},
		OTHER: UserGroupGroupSource{
			value: "OTHER",
		},
	}
}

func (c UserGroupGroupSource) Value() string {
	return c.value
}

func (c UserGroupGroupSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserGroupGroupSource) UnmarshalJSON(b []byte) error {
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
