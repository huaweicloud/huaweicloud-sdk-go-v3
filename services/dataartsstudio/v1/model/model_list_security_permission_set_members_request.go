package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityPermissionSetMembersRequest Request Object
type ListSecurityPermissionSetMembersRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 成员名称
	MemberName *string `json:"member_name,omitempty"`

	// 成员类型,USER,USER_GROUP,WORKSPACE_ROLE
	MemberType *ListSecurityPermissionSetMembersRequestMemberType `json:"member_type,omitempty"`

	// 是否升序（仅指定排序参数时有效）
	OrderByAsc *bool `json:"order_by_asc,omitempty"`

	// 排序参数, CREATE_TIME, MEMBER_NAME
	OrderBy *ListSecurityPermissionSetMembersRequestOrderBy `json:"order_by,omitempty"`
}

func (o ListSecurityPermissionSetMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPermissionSetMembersRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityPermissionSetMembersRequest", string(data)}, " ")
}

type ListSecurityPermissionSetMembersRequestMemberType struct {
	value string
}

type ListSecurityPermissionSetMembersRequestMemberTypeEnum struct {
	USER           ListSecurityPermissionSetMembersRequestMemberType
	USER_GROUP     ListSecurityPermissionSetMembersRequestMemberType
	WORKSPACE_ROLE ListSecurityPermissionSetMembersRequestMemberType
}

func GetListSecurityPermissionSetMembersRequestMemberTypeEnum() ListSecurityPermissionSetMembersRequestMemberTypeEnum {
	return ListSecurityPermissionSetMembersRequestMemberTypeEnum{
		USER: ListSecurityPermissionSetMembersRequestMemberType{
			value: "USER",
		},
		USER_GROUP: ListSecurityPermissionSetMembersRequestMemberType{
			value: "USER_GROUP",
		},
		WORKSPACE_ROLE: ListSecurityPermissionSetMembersRequestMemberType{
			value: "WORKSPACE_ROLE",
		},
	}
}

func (c ListSecurityPermissionSetMembersRequestMemberType) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetMembersRequestMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetMembersRequestMemberType) UnmarshalJSON(b []byte) error {
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

type ListSecurityPermissionSetMembersRequestOrderBy struct {
	value string
}

type ListSecurityPermissionSetMembersRequestOrderByEnum struct {
	CREATE_TIME ListSecurityPermissionSetMembersRequestOrderBy
	MEMBER_NAME ListSecurityPermissionSetMembersRequestOrderBy
}

func GetListSecurityPermissionSetMembersRequestOrderByEnum() ListSecurityPermissionSetMembersRequestOrderByEnum {
	return ListSecurityPermissionSetMembersRequestOrderByEnum{
		CREATE_TIME: ListSecurityPermissionSetMembersRequestOrderBy{
			value: "CREATE_TIME",
		},
		MEMBER_NAME: ListSecurityPermissionSetMembersRequestOrderBy{
			value: "MEMBER_NAME",
		},
	}
}

func (c ListSecurityPermissionSetMembersRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetMembersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetMembersRequestOrderBy) UnmarshalJSON(b []byte) error {
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
