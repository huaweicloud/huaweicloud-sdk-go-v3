package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionSetMemberCreateDto struct {

	// 成员类型, 用户/用户组/工作空间角色, USER, USER_GROUP, WORKSPACE_ROLE
	MemberType *PermissionSetMemberCreateDtoMemberType `json:"member_type,omitempty"`

	// 成员id
	MemberId *string `json:"member_id,omitempty"`

	// 成员名称
	MemberName *string `json:"member_name,omitempty"`

	// 工作空间
	Workspace *string `json:"workspace,omitempty"`
}

func (o PermissionSetMemberCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSetMemberCreateDto struct{}"
	}

	return strings.Join([]string{"PermissionSetMemberCreateDto", string(data)}, " ")
}

type PermissionSetMemberCreateDtoMemberType struct {
	value string
}

type PermissionSetMemberCreateDtoMemberTypeEnum struct {
	USER           PermissionSetMemberCreateDtoMemberType
	USER_GROUP     PermissionSetMemberCreateDtoMemberType
	WORKSPACE_ROLE PermissionSetMemberCreateDtoMemberType
}

func GetPermissionSetMemberCreateDtoMemberTypeEnum() PermissionSetMemberCreateDtoMemberTypeEnum {
	return PermissionSetMemberCreateDtoMemberTypeEnum{
		USER: PermissionSetMemberCreateDtoMemberType{
			value: "USER",
		},
		USER_GROUP: PermissionSetMemberCreateDtoMemberType{
			value: "USER_GROUP",
		},
		WORKSPACE_ROLE: PermissionSetMemberCreateDtoMemberType{
			value: "WORKSPACE_ROLE",
		},
	}
}

func (c PermissionSetMemberCreateDtoMemberType) Value() string {
	return c.value
}

func (c PermissionSetMemberCreateDtoMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetMemberCreateDtoMemberType) UnmarshalJSON(b []byte) error {
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
