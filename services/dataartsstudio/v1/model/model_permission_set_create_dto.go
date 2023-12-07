package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionSetCreateDto struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 父权限集id
	ParentId *string `json:"parent_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 权限集类型, COMMON, MRS_MANAGED
	Type *PermissionSetCreateDtoType `json:"type,omitempty"`

	// 纳管角色所在集群id（仅纳管类权限集需要）
	ManagedClusterId *string `json:"managed_cluster_id,omitempty"`

	// 纳管角色所在集群名称（仅纳管类权限集需要）
	ManagedClusterName *string `json:"managed_cluster_name,omitempty"`

	// 纳管角色名称（仅纳管类权限集需要）
	ManagedRoleName *string `json:"managed_role_name,omitempty"`

	// 管理员id
	ManagerId *string `json:"manager_id,omitempty"`

	// 管理员名称
	ManagerName *string `json:"manager_name,omitempty"`

	// 管理员类型, 用户/用户组, USER, USER_GROUP
	ManagerType *PermissionSetCreateDtoManagerType `json:"manager_type,omitempty"`
}

func (o PermissionSetCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSetCreateDto struct{}"
	}

	return strings.Join([]string{"PermissionSetCreateDto", string(data)}, " ")
}

type PermissionSetCreateDtoType struct {
	value string
}

type PermissionSetCreateDtoTypeEnum struct {
	COMMON      PermissionSetCreateDtoType
	MRS_MANAGED PermissionSetCreateDtoType
}

func GetPermissionSetCreateDtoTypeEnum() PermissionSetCreateDtoTypeEnum {
	return PermissionSetCreateDtoTypeEnum{
		COMMON: PermissionSetCreateDtoType{
			value: "COMMON",
		},
		MRS_MANAGED: PermissionSetCreateDtoType{
			value: "MRS_MANAGED",
		},
	}
}

func (c PermissionSetCreateDtoType) Value() string {
	return c.value
}

func (c PermissionSetCreateDtoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetCreateDtoType) UnmarshalJSON(b []byte) error {
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

type PermissionSetCreateDtoManagerType struct {
	value string
}

type PermissionSetCreateDtoManagerTypeEnum struct {
	USER       PermissionSetCreateDtoManagerType
	USER_GROUP PermissionSetCreateDtoManagerType
}

func GetPermissionSetCreateDtoManagerTypeEnum() PermissionSetCreateDtoManagerTypeEnum {
	return PermissionSetCreateDtoManagerTypeEnum{
		USER: PermissionSetCreateDtoManagerType{
			value: "USER",
		},
		USER_GROUP: PermissionSetCreateDtoManagerType{
			value: "USER_GROUP",
		},
	}
}

func (c PermissionSetCreateDtoManagerType) Value() string {
	return c.value
}

func (c PermissionSetCreateDtoManagerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetCreateDtoManagerType) UnmarshalJSON(b []byte) error {
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
