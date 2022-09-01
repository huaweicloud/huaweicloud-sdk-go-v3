package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type User struct {

	// 用户ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用权限角色 - read：应用下资源只读权限，至少要存在此权限，包括API调试 - access：应用下资源的访问管理权限 - delete：应用下资源的删除权限 - modify：应用下资源的修改权限，包括API发布、下线 - admin：应用和应用下资源的权限 - 仅提供admin时，会自动应用其它所有权限 - 未提供read时会自动应用read权限
	Roles *[]UserRoles `json:"roles,omitempty" xml:"roles"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}

type UserRoles struct {
	value string
}

type UserRolesEnum struct {
	READ   UserRoles
	ACCESS UserRoles
	DELETE UserRoles
	MODIFY UserRoles
	ADMIN  UserRoles
}

func GetUserRolesEnum() UserRolesEnum {
	return UserRolesEnum{
		READ: UserRoles{
			value: "read",
		},
		ACCESS: UserRoles{
			value: "access",
		},
		DELETE: UserRoles{
			value: "delete",
		},
		MODIFY: UserRoles{
			value: "modify",
		},
		ADMIN: UserRoles{
			value: "admin",
		},
	}
}

func (c UserRoles) Value() string {
	return c.value
}

func (c UserRoles) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserRoles) UnmarshalJSON(b []byte) error {
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
