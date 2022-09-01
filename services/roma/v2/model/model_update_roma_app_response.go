package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateRomaAppResponse struct {

	// 应用ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用名称 - 字符集：支持中文、英文字母、数字、中划线、下划线、点、空格和中英文圆括号 - 约束：实例下唯一
	Name *string `json:"name,omitempty" xml:"name"`

	// 应用描述
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 应用权限角色 - read：应用下资源只读权限，至少要存在此权限，包括API调试 - access：应用下资源的访问管理权限 - delete：应用下资源的删除权限 - modify：应用下资源的修改权限，包括API发布、下线 - admin：应用和应用下资源的权限 - 仅提供admin时，会自动应用其它所有权限 - 未提供read时会自动应用read权限
	Roles *[]UpdateRomaAppResponseRoles `json:"roles,omitempty" xml:"roles"`

	// 创建UTC时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 创建UTC时间
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	CreatedUser *ServerAppInfoCreatedUser `json:"created_user,omitempty" xml:"created_user"`

	LastUpdatedUser *ServerAppInfoLastUpdatedUser `json:"last_updated_user,omitempty" xml:"last_updated_user"`

	// 是否是应用拥有者
	Owner *bool `json:"owner,omitempty" xml:"owner"`

	// 应用认证访问KEY,未提供时随机生成 - 字符集：支持中文、英文字母、数字、中划线、下划线、@号和点，以字母或中文或数字开头 - 约束：实例下唯一
	Key *string `json:"key,omitempty" xml:"key"`

	// 是否收藏应用，收藏的应用会在列表里优先显示
	Favorite       *bool `json:"favorite,omitempty" xml:"favorite"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateRomaAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRomaAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateRomaAppResponse", string(data)}, " ")
}

type UpdateRomaAppResponseRoles struct {
	value string
}

type UpdateRomaAppResponseRolesEnum struct {
	READ   UpdateRomaAppResponseRoles
	ACCESS UpdateRomaAppResponseRoles
	DELETE UpdateRomaAppResponseRoles
	MODIFY UpdateRomaAppResponseRoles
	ADMIN  UpdateRomaAppResponseRoles
}

func GetUpdateRomaAppResponseRolesEnum() UpdateRomaAppResponseRolesEnum {
	return UpdateRomaAppResponseRolesEnum{
		READ: UpdateRomaAppResponseRoles{
			value: "read",
		},
		ACCESS: UpdateRomaAppResponseRoles{
			value: "access",
		},
		DELETE: UpdateRomaAppResponseRoles{
			value: "delete",
		},
		MODIFY: UpdateRomaAppResponseRoles{
			value: "modify",
		},
		ADMIN: UpdateRomaAppResponseRoles{
			value: "admin",
		},
	}
}

func (c UpdateRomaAppResponseRoles) Value() string {
	return c.value
}

func (c UpdateRomaAppResponseRoles) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRomaAppResponseRoles) UnmarshalJSON(b []byte) error {
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
