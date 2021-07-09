package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type CreateUsersInfo struct {
	// DDM实例帐号名称，命名要求如下。  - 长度为6-32个字符。 - 必须以字母开头。 - 可以包含字母，数字、下划线，不能包含其它特殊字符。

	Name string `json:"name"`
	// DDM实例帐号密码。

	Password string `json:"password"`
	// DDM实例帐号的基础权限。  取值为：CREATE、DROP、ALTER、INDEX、INSERT、DELETE、UPDATE、SELECT

	BaseAuthority []CreateUsersInfoBaseAuthority `json:"base_authority"`
	// DDM实例帐号的扩展权限，默认值为空。  取值范围为：fulltableDelete、fulltableSelect、fulltableUpdate。  说明： 权限配置应该遵循如下原则：  请至少选择一个基础权限，且扩展权限对应的基础权限必须选择，对应关系如下：   - “fulltableSelect”对应“SELECT”   - “fulltableDelete”对应“DELETE”   - “fulltableUpdate”对应“UPDATE”

	ExtendAuthority *[]CreateUsersInfoExtendAuthority `json:"extend_authority,omitempty"`
	// DDM实例帐号的描述，最大长度不能超过256。默认值为空。

	Description *string `json:"description,omitempty"`
	// 关联的逻辑库的集合。 databases字段可以省略，即创建用户时可以不关联逻辑库。

	Databases *[]CreateUsersDatabases `json:"databases,omitempty"`
}

func (o CreateUsersInfo) String() string {
	data, err := json.Marshal(o)
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
	return json.Marshal(c.value)
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

type CreateUsersInfoExtendAuthority struct {
	value string
}

type CreateUsersInfoExtendAuthorityEnum struct {
	FULLTABLE_SELECT CreateUsersInfoExtendAuthority
	FULLTABLE_DELETE CreateUsersInfoExtendAuthority
	FULLTABLE_UPDATE CreateUsersInfoExtendAuthority
}

func GetCreateUsersInfoExtendAuthorityEnum() CreateUsersInfoExtendAuthorityEnum {
	return CreateUsersInfoExtendAuthorityEnum{
		FULLTABLE_SELECT: CreateUsersInfoExtendAuthority{
			value: "fulltableSelect",
		},
		FULLTABLE_DELETE: CreateUsersInfoExtendAuthority{
			value: "fulltableDelete",
		},
		FULLTABLE_UPDATE: CreateUsersInfoExtendAuthority{
			value: "fulltableUpdate",
		},
	}
}

func (c CreateUsersInfoExtendAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateUsersInfoExtendAuthority) UnmarshalJSON(b []byte) error {
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
