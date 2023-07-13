package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AppUsersUsers struct {

	// 用户ID
	Id *string `json:"id,omitempty"`

	// 用户名称
	Name *string `json:"name,omitempty"`

	// 应用权限角色 - read：应用下资源只读权限，至少要存在此权限，包括API调试 - access：应用下资源的访问管理权限 - delete：应用下资源的删除权限 - modify：应用下资源的修改权限，包括API发布、下线 - admin：应用和应用下资源的权限 - 仅提供admin时，会自动应用其它所有权限 - 未提供read时会自动应用read权限
	Roles *[]AppUsersUsersRoles `json:"roles,omitempty"`
}

func (o AppUsersUsers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppUsersUsers struct{}"
	}

	return strings.Join([]string{"AppUsersUsers", string(data)}, " ")
}

type AppUsersUsersRoles struct {
	value string
}

type AppUsersUsersRolesEnum struct {
	READ   AppUsersUsersRoles
	ACCESS AppUsersUsersRoles
	DELETE AppUsersUsersRoles
	MODIFY AppUsersUsersRoles
	ADMIN  AppUsersUsersRoles
}

func GetAppUsersUsersRolesEnum() AppUsersUsersRolesEnum {
	return AppUsersUsersRolesEnum{
		READ: AppUsersUsersRoles{
			value: "read",
		},
		ACCESS: AppUsersUsersRoles{
			value: "access",
		},
		DELETE: AppUsersUsersRoles{
			value: "delete",
		},
		MODIFY: AppUsersUsersRoles{
			value: "modify",
		},
		ADMIN: AppUsersUsersRoles{
			value: "admin",
		},
	}
}

func (c AppUsersUsersRoles) Value() string {
	return c.value
}

func (c AppUsersUsersRoles) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppUsersUsersRoles) UnmarshalJSON(b []byte) error {
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
