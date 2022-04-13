package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// This is a auto create Body Object
type CreateUsersInfo struct {
	// DDM实例帐号名称，命名要求如下。  - 长度为1-32个字符。 - 必须以字母开头。 - 可以包含字母，数字、下划线，不能包含其它特殊字符。

	Name string `json:"name"`
	// DDM实例帐号密码。

	Password string `json:"password"`
	// DDM实例帐号的基础权限。  取值为：CREATE、DROP、ALTER、INDEX、INSERT、DELETE、UPDATE、SELECT

	BaseAuthority []CreateUsersInfoBaseAuthority `json:"base_authority"`
	// DDM实例帐号的描述，最大长度不能超过256。默认值为空。

	Description *string `json:"description,omitempty"`
	// 关联的逻辑库的集合。 databases字段可以省略，即创建用户时可以不关联逻辑库。

	Databases *[]CreateUsersDatabases `json:"databases,omitempty"`
}

func (o CreateUsersInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUsersInfo struct{}"
	}

	return strings.Join([]string{"CreateUsersInfo", string(data)}, " ")
}

type CreateUsersInfoBaseAuthority struct {
	value string
}

type CreateUsersInfoBaseAuthorityEnum struct {
	CREATE CreateUsersInfoBaseAuthority
	DROP   CreateUsersInfoBaseAuthority
	ALTER  CreateUsersInfoBaseAuthority
	INDEX  CreateUsersInfoBaseAuthority
	INSERT CreateUsersInfoBaseAuthority
	DELETE CreateUsersInfoBaseAuthority
	UPDATE CreateUsersInfoBaseAuthority
	SELECT CreateUsersInfoBaseAuthority
}

func GetCreateUsersInfoBaseAuthorityEnum() CreateUsersInfoBaseAuthorityEnum {
	return CreateUsersInfoBaseAuthorityEnum{
		CREATE: CreateUsersInfoBaseAuthority{
			value: "CREATE",
		},
		DROP: CreateUsersInfoBaseAuthority{
			value: "DROP",
		},
		ALTER: CreateUsersInfoBaseAuthority{
			value: "ALTER",
		},
		INDEX: CreateUsersInfoBaseAuthority{
			value: "INDEX",
		},
		INSERT: CreateUsersInfoBaseAuthority{
			value: "INSERT",
		},
		DELETE: CreateUsersInfoBaseAuthority{
			value: "DELETE",
		},
		UPDATE: CreateUsersInfoBaseAuthority{
			value: "UPDATE",
		},
		SELECT: CreateUsersInfoBaseAuthority{
			value: "SELECT",
		},
	}
}

func (c CreateUsersInfoBaseAuthority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateUsersInfoBaseAuthority) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
