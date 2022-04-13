package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type CreateRomaAppResponse struct {
	// 应用ID

	Id *string `json:"id,omitempty"`
	// 应用名称 - 字符集：支持中文、英文字母、数字、中划线、下划线、点、空格和中英文圆括号 - 约束：实例下唯一

	Name *string `json:"name,omitempty"`
	// 应用描述

	Remark *string `json:"remark,omitempty"`
	// 应用权限角色 - read：应用下资源只读权限，至少要存在此权限，包括API调试 - access：应用下资源的访问管理权限 - delete：应用下资源的删除权限 - modify：应用下资源的修改权限，包括API发布、下线 - admin：应用和应用下资源的权限 - 仅提供admin时，会自动应用其它所有权限 - 未提供read时会自动应用read权限

	Roles *[]CreateRomaAppResponseRoles `json:"roles,omitempty"`
	// 创建UTC时间

	CreateTime *string `json:"create_time,omitempty"`
	// 创建UTC时间

	UpdateTime *string `json:"update_time,omitempty"`

	CreatedUser *ServerAppInfoCreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *ServerAppInfoLastUpdatedUser `json:"last_updated_user,omitempty"`
	// 是否是应用拥有者

	Owner *bool `json:"owner,omitempty"`
	// 应用认证访问KEY,未提供时随机生成 - 字符集：支持中文、英文字母、数字、中划线、下划线、@号和点，以字母或中文或数字开头 - 约束：实例下唯一

	Key *string `json:"key,omitempty"`
	// 是否收藏应用，收藏的应用会在列表里优先显示

	Favorite       *bool `json:"favorite,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateRomaAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRomaAppResponse struct{}"
	}

	return strings.Join([]string{"CreateRomaAppResponse", string(data)}, " ")
}

type CreateRomaAppResponseRoles struct {
	value string
}

type CreateRomaAppResponseRolesEnum struct {
	READ   CreateRomaAppResponseRoles
	ACCESS CreateRomaAppResponseRoles
	DELETE CreateRomaAppResponseRoles
	MODIFY CreateRomaAppResponseRoles
	ADMIN  CreateRomaAppResponseRoles
}

func GetCreateRomaAppResponseRolesEnum() CreateRomaAppResponseRolesEnum {
	return CreateRomaAppResponseRolesEnum{
		READ: CreateRomaAppResponseRoles{
			value: "read",
		},
		ACCESS: CreateRomaAppResponseRoles{
			value: "access",
		},
		DELETE: CreateRomaAppResponseRoles{
			value: "delete",
		},
		MODIFY: CreateRomaAppResponseRoles{
			value: "modify",
		},
		ADMIN: CreateRomaAppResponseRoles{
			value: "admin",
		},
	}
}

func (c CreateRomaAppResponseRoles) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRomaAppResponseRoles) UnmarshalJSON(b []byte) error {
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
