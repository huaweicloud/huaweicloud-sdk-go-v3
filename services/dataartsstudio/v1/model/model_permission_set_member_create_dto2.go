package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionSetMemberCreateDto2 struct {

	// 成员类型, 用户/用户组/工作空间角色, USER, USER_GROUP, WORKSPACE_ROLE
	MemberType *PermissionSetMemberCreateDto2MemberType `json:"member_type,omitempty"`

	// 成员id
	MemberId *string `json:"member_id,omitempty"`

	// 成员名称
	MemberName *string `json:"member_name,omitempty"`

	// 工作空间
	Workspace *string `json:"workspace,omitempty"`
}

func (o PermissionSetMemberCreateDto2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSetMemberCreateDto2 struct{}"
	}

	return strings.Join([]string{"PermissionSetMemberCreateDto2", string(data)}, " ")
}

type PermissionSetMemberCreateDto2MemberType struct {
	value string
}

type PermissionSetMemberCreateDto2MemberTypeEnum struct {
	USER           PermissionSetMemberCreateDto2MemberType
	USER_GROUP     PermissionSetMemberCreateDto2MemberType
	WORKSPACE_ROLE PermissionSetMemberCreateDto2MemberType
}

func GetPermissionSetMemberCreateDto2MemberTypeEnum() PermissionSetMemberCreateDto2MemberTypeEnum {
	return PermissionSetMemberCreateDto2MemberTypeEnum{
		USER: PermissionSetMemberCreateDto2MemberType{
			value: "USER",
		},
		USER_GROUP: PermissionSetMemberCreateDto2MemberType{
			value: "USER_GROUP",
		},
		WORKSPACE_ROLE: PermissionSetMemberCreateDto2MemberType{
			value: "WORKSPACE_ROLE",
		},
	}
}

func (c PermissionSetMemberCreateDto2MemberType) Value() string {
	return c.value
}

func (c PermissionSetMemberCreateDto2MemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetMemberCreateDto2MemberType) UnmarshalJSON(b []byte) error {
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
