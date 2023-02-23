package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// access policy input body
type AccessPolicyInput struct {

	// 主体信息
	PrincipalList []Principal `json:"principal_list"`

	Resource *ResourceInfo `json:"resource"`

	// 拒绝/允许
	Effect bool `json:"effect"`

	// 权限列表
	Permissions []AccessPolicyInputPermissions `json:"permissions"`

	// 可传递的权限列表
	GrantAblePermissions *[]AccessPolicyInputGrantAblePermissions `json:"grant_able_permissions,omitempty"`

	// 条件
	Conditions *string `json:"conditions,omitempty"`

	// 行过滤
	DataFilter *string `json:"data_filter,omitempty"`

	// 列掩码
	DataMask *string `json:"data_mask,omitempty"`
}

func (o AccessPolicyInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicyInput struct{}"
	}

	return strings.Join([]string{"AccessPolicyInput", string(data)}, " ")
}

type AccessPolicyInputPermissions struct {
	value string
}

type AccessPolicyInputPermissionsEnum struct {
	ALL             AccessPolicyInputPermissions
	CREATE          AccessPolicyInputPermissions
	ALTER           AccessPolicyInputPermissions
	DROP            AccessPolicyInputPermissions
	DESCRIBE        AccessPolicyInputPermissions
	EXEC            AccessPolicyInputPermissions
	CREATE_DATABASE AccessPolicyInputPermissions
	LIST_DATABASE   AccessPolicyInputPermissions
	CREATE_TABLE    AccessPolicyInputPermissions
	LIST_TABLE      AccessPolicyInputPermissions
	CREATE_FUNC     AccessPolicyInputPermissions
	LIST_FUNC       AccessPolicyInputPermissions
	REGISTER_MODEL  AccessPolicyInputPermissions
	LIST_MODEL      AccessPolicyInputPermissions
	INSERT          AccessPolicyInputPermissions
	UPDATE          AccessPolicyInputPermissions
	DELETE          AccessPolicyInputPermissions
	SELECT          AccessPolicyInputPermissions
	READ            AccessPolicyInputPermissions
	WRITE           AccessPolicyInputPermissions
	OPERATE         AccessPolicyInputPermissions
}

func GetAccessPolicyInputPermissionsEnum() AccessPolicyInputPermissionsEnum {
	return AccessPolicyInputPermissionsEnum{
		ALL: AccessPolicyInputPermissions{
			value: "ALL",
		},
		CREATE: AccessPolicyInputPermissions{
			value: "CREATE",
		},
		ALTER: AccessPolicyInputPermissions{
			value: "ALTER",
		},
		DROP: AccessPolicyInputPermissions{
			value: "DROP",
		},
		DESCRIBE: AccessPolicyInputPermissions{
			value: "DESCRIBE",
		},
		EXEC: AccessPolicyInputPermissions{
			value: "EXEC",
		},
		CREATE_DATABASE: AccessPolicyInputPermissions{
			value: "CREATE_DATABASE",
		},
		LIST_DATABASE: AccessPolicyInputPermissions{
			value: "LIST_DATABASE",
		},
		CREATE_TABLE: AccessPolicyInputPermissions{
			value: "CREATE_TABLE",
		},
		LIST_TABLE: AccessPolicyInputPermissions{
			value: "LIST_TABLE",
		},
		CREATE_FUNC: AccessPolicyInputPermissions{
			value: "CREATE_FUNC",
		},
		LIST_FUNC: AccessPolicyInputPermissions{
			value: "LIST_FUNC",
		},
		REGISTER_MODEL: AccessPolicyInputPermissions{
			value: "REGISTER_MODEL",
		},
		LIST_MODEL: AccessPolicyInputPermissions{
			value: "LIST_MODEL",
		},
		INSERT: AccessPolicyInputPermissions{
			value: "INSERT",
		},
		UPDATE: AccessPolicyInputPermissions{
			value: "UPDATE",
		},
		DELETE: AccessPolicyInputPermissions{
			value: "DELETE",
		},
		SELECT: AccessPolicyInputPermissions{
			value: "SELECT",
		},
		READ: AccessPolicyInputPermissions{
			value: "READ",
		},
		WRITE: AccessPolicyInputPermissions{
			value: "WRITE",
		},
		OPERATE: AccessPolicyInputPermissions{
			value: "OPERATE",
		},
	}
}

func (c AccessPolicyInputPermissions) Value() string {
	return c.value
}

func (c AccessPolicyInputPermissions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPolicyInputPermissions) UnmarshalJSON(b []byte) error {
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

type AccessPolicyInputGrantAblePermissions struct {
	value string
}

type AccessPolicyInputGrantAblePermissionsEnum struct {
	ALL             AccessPolicyInputGrantAblePermissions
	CREATE          AccessPolicyInputGrantAblePermissions
	ALTER           AccessPolicyInputGrantAblePermissions
	DROP            AccessPolicyInputGrantAblePermissions
	DESCRIBE        AccessPolicyInputGrantAblePermissions
	EXEC            AccessPolicyInputGrantAblePermissions
	CREATE_DATABASE AccessPolicyInputGrantAblePermissions
	LIST_DATABASE   AccessPolicyInputGrantAblePermissions
	CREATE_TABLE    AccessPolicyInputGrantAblePermissions
	LIST_TABLE      AccessPolicyInputGrantAblePermissions
	CREATE_FUNC     AccessPolicyInputGrantAblePermissions
	LIST_FUNC       AccessPolicyInputGrantAblePermissions
	REGISTER_MODEL  AccessPolicyInputGrantAblePermissions
	LIST_MODEL      AccessPolicyInputGrantAblePermissions
	INSERT          AccessPolicyInputGrantAblePermissions
	UPDATE          AccessPolicyInputGrantAblePermissions
	DELETE          AccessPolicyInputGrantAblePermissions
	SELECT          AccessPolicyInputGrantAblePermissions
	READ            AccessPolicyInputGrantAblePermissions
	WRITE           AccessPolicyInputGrantAblePermissions
	OPERATE         AccessPolicyInputGrantAblePermissions
}

func GetAccessPolicyInputGrantAblePermissionsEnum() AccessPolicyInputGrantAblePermissionsEnum {
	return AccessPolicyInputGrantAblePermissionsEnum{
		ALL: AccessPolicyInputGrantAblePermissions{
			value: "ALL",
		},
		CREATE: AccessPolicyInputGrantAblePermissions{
			value: "CREATE",
		},
		ALTER: AccessPolicyInputGrantAblePermissions{
			value: "ALTER",
		},
		DROP: AccessPolicyInputGrantAblePermissions{
			value: "DROP",
		},
		DESCRIBE: AccessPolicyInputGrantAblePermissions{
			value: "DESCRIBE",
		},
		EXEC: AccessPolicyInputGrantAblePermissions{
			value: "EXEC",
		},
		CREATE_DATABASE: AccessPolicyInputGrantAblePermissions{
			value: "CREATE_DATABASE",
		},
		LIST_DATABASE: AccessPolicyInputGrantAblePermissions{
			value: "LIST_DATABASE",
		},
		CREATE_TABLE: AccessPolicyInputGrantAblePermissions{
			value: "CREATE_TABLE",
		},
		LIST_TABLE: AccessPolicyInputGrantAblePermissions{
			value: "LIST_TABLE",
		},
		CREATE_FUNC: AccessPolicyInputGrantAblePermissions{
			value: "CREATE_FUNC",
		},
		LIST_FUNC: AccessPolicyInputGrantAblePermissions{
			value: "LIST_FUNC",
		},
		REGISTER_MODEL: AccessPolicyInputGrantAblePermissions{
			value: "REGISTER_MODEL",
		},
		LIST_MODEL: AccessPolicyInputGrantAblePermissions{
			value: "LIST_MODEL",
		},
		INSERT: AccessPolicyInputGrantAblePermissions{
			value: "INSERT",
		},
		UPDATE: AccessPolicyInputGrantAblePermissions{
			value: "UPDATE",
		},
		DELETE: AccessPolicyInputGrantAblePermissions{
			value: "DELETE",
		},
		SELECT: AccessPolicyInputGrantAblePermissions{
			value: "SELECT",
		},
		READ: AccessPolicyInputGrantAblePermissions{
			value: "READ",
		},
		WRITE: AccessPolicyInputGrantAblePermissions{
			value: "WRITE",
		},
		OPERATE: AccessPolicyInputGrantAblePermissions{
			value: "OPERATE",
		},
	}
}

func (c AccessPolicyInputGrantAblePermissions) Value() string {
	return c.value
}

func (c AccessPolicyInputGrantAblePermissions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPolicyInputGrantAblePermissions) UnmarshalJSON(b []byte) error {
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
