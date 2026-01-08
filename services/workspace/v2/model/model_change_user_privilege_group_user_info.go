package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeUserPrivilegeGroupUserInfo 桌面分配用户信息。
type ChangeUserPrivilegeGroupUserInfo struct {

	// 桌面分配对象的名称，当type类型USER时填写用户名字；当type类型GROUP时填写用户组名称。
	UserName string `json:"user_name"`

	// 用户所属域。
	Domain *string `json:"domain,omitempty"`

	// 桌面用户所属的用户组。 - sudo：Linux管理员组。 - default：Linux默认用户组。 - administrators：Windows管理员组。管理员拥有对该桌面的完全访问权，可以做任何需要的更改（禁用操作除外）。 - users：Windows标准用户组。标准用户可以使用大多数软件，并可以更改不影响其他用户的系统设置。
	UserPrivilegeGroup *string `json:"user_privilege_group,omitempty"`

	// 对象类型，可选值为： - USER：用户。 - GROUP：用户组。
	Type ChangeUserPrivilegeGroupUserInfoType `json:"type"`
}

func (o ChangeUserPrivilegeGroupUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeUserPrivilegeGroupUserInfo struct{}"
	}

	return strings.Join([]string{"ChangeUserPrivilegeGroupUserInfo", string(data)}, " ")
}

type ChangeUserPrivilegeGroupUserInfoType struct {
	value string
}

type ChangeUserPrivilegeGroupUserInfoTypeEnum struct {
	USER  ChangeUserPrivilegeGroupUserInfoType
	GROUP ChangeUserPrivilegeGroupUserInfoType
}

func GetChangeUserPrivilegeGroupUserInfoTypeEnum() ChangeUserPrivilegeGroupUserInfoTypeEnum {
	return ChangeUserPrivilegeGroupUserInfoTypeEnum{
		USER: ChangeUserPrivilegeGroupUserInfoType{
			value: "USER",
		},
		GROUP: ChangeUserPrivilegeGroupUserInfoType{
			value: "GROUP",
		},
	}
}

func (c ChangeUserPrivilegeGroupUserInfoType) Value() string {
	return c.value
}

func (c ChangeUserPrivilegeGroupUserInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeUserPrivilegeGroupUserInfoType) UnmarshalJSON(b []byte) error {
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
