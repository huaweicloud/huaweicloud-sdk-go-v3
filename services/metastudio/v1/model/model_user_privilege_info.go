package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 增加租户权限
type UserPrivilegeInfo struct {

	// 租户ID
	TenantId string `json:"tenant_id"`

	// 租户权限 INTERNAL_TEST: 内测用户权限，有服务配额限制 SYSTEM_ADMIN：系统管理员权限，可加权限 PARTNER：伙伴权限，暂不做配额限制 ME_PRIVILEGE:ME白名单权限
	Privilege UserPrivilegeInfoPrivilege `json:"privilege"`

	// 配额过期时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o UserPrivilegeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserPrivilegeInfo struct{}"
	}

	return strings.Join([]string{"UserPrivilegeInfo", string(data)}, " ")
}

type UserPrivilegeInfoPrivilege struct {
	value string
}

type UserPrivilegeInfoPrivilegeEnum struct {
	INTERNAL_TEST UserPrivilegeInfoPrivilege
	SYSTEM_ADMIN  UserPrivilegeInfoPrivilege
	PARTNER       UserPrivilegeInfoPrivilege
	ME_PRIVILEGE  UserPrivilegeInfoPrivilege
}

func GetUserPrivilegeInfoPrivilegeEnum() UserPrivilegeInfoPrivilegeEnum {
	return UserPrivilegeInfoPrivilegeEnum{
		INTERNAL_TEST: UserPrivilegeInfoPrivilege{
			value: "INTERNAL_TEST",
		},
		SYSTEM_ADMIN: UserPrivilegeInfoPrivilege{
			value: "SYSTEM_ADMIN",
		},
		PARTNER: UserPrivilegeInfoPrivilege{
			value: "PARTNER",
		},
		ME_PRIVILEGE: UserPrivilegeInfoPrivilege{
			value: "ME_PRIVILEGE",
		},
	}
}

func (c UserPrivilegeInfoPrivilege) Value() string {
	return c.value
}

func (c UserPrivilegeInfoPrivilege) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserPrivilegeInfoPrivilege) UnmarshalJSON(b []byte) error {
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
