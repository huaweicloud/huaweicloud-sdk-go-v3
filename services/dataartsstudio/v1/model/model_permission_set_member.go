package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionSetMember struct {

	// id
	Id *string `json:"id,omitempty"`

	// 权限集id
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 成员类型, 用户/用户组/工作空间角色(废弃)/集群角色, USER, USER_GROUP, WORKSPACE_ROLE, CLUSTER_ROLE
	MemberType *PermissionSetMemberMemberType `json:"member_type,omitempty"`

	// 成员id
	MemberId *string `json:"member_id,omitempty"`

	// 成员名称
	MemberName *string `json:"member_name,omitempty"`

	// 成员状态, NORMAL, UNFINISHED
	MemberStatus *PermissionSetMemberMemberStatus `json:"member_status,omitempty"`

	// 工作空间(仅工作空间角色需要)
	Workspace *string `json:"workspace,omitempty"`

	// 集群类型(仅集群角色需要), MRS, DWS, DLI
	ClusterType *PermissionSetMemberClusterType `json:"cluster_type,omitempty"`

	// 集群id(仅集群角色需要)
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称(仅集群角色需要)
	ClusterName *string `json:"cluster_name,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 到期时间
	Deadline *int64 `json:"deadline,omitempty"`
}

func (o PermissionSetMember) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSetMember struct{}"
	}

	return strings.Join([]string{"PermissionSetMember", string(data)}, " ")
}

type PermissionSetMemberMemberType struct {
	value string
}

type PermissionSetMemberMemberTypeEnum struct {
	USER           PermissionSetMemberMemberType
	USER_GROUP     PermissionSetMemberMemberType
	WORKSPACE_ROLE PermissionSetMemberMemberType
	CLUSTER_ROLE   PermissionSetMemberMemberType
}

func GetPermissionSetMemberMemberTypeEnum() PermissionSetMemberMemberTypeEnum {
	return PermissionSetMemberMemberTypeEnum{
		USER: PermissionSetMemberMemberType{
			value: "USER",
		},
		USER_GROUP: PermissionSetMemberMemberType{
			value: "USER_GROUP",
		},
		WORKSPACE_ROLE: PermissionSetMemberMemberType{
			value: "WORKSPACE_ROLE",
		},
		CLUSTER_ROLE: PermissionSetMemberMemberType{
			value: "CLUSTER_ROLE",
		},
	}
}

func (c PermissionSetMemberMemberType) Value() string {
	return c.value
}

func (c PermissionSetMemberMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetMemberMemberType) UnmarshalJSON(b []byte) error {
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

type PermissionSetMemberMemberStatus struct {
	value string
}

type PermissionSetMemberMemberStatusEnum struct {
	NORMAL     PermissionSetMemberMemberStatus
	UNFINISHED PermissionSetMemberMemberStatus
}

func GetPermissionSetMemberMemberStatusEnum() PermissionSetMemberMemberStatusEnum {
	return PermissionSetMemberMemberStatusEnum{
		NORMAL: PermissionSetMemberMemberStatus{
			value: "NORMAL",
		},
		UNFINISHED: PermissionSetMemberMemberStatus{
			value: "UNFINISHED",
		},
	}
}

func (c PermissionSetMemberMemberStatus) Value() string {
	return c.value
}

func (c PermissionSetMemberMemberStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetMemberMemberStatus) UnmarshalJSON(b []byte) error {
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

type PermissionSetMemberClusterType struct {
	value string
}

type PermissionSetMemberClusterTypeEnum struct {
	MRS PermissionSetMemberClusterType
	DWS PermissionSetMemberClusterType
	DLI PermissionSetMemberClusterType
}

func GetPermissionSetMemberClusterTypeEnum() PermissionSetMemberClusterTypeEnum {
	return PermissionSetMemberClusterTypeEnum{
		MRS: PermissionSetMemberClusterType{
			value: "MRS",
		},
		DWS: PermissionSetMemberClusterType{
			value: "DWS",
		},
		DLI: PermissionSetMemberClusterType{
			value: "DLI",
		},
	}
}

func (c PermissionSetMemberClusterType) Value() string {
	return c.value
}

func (c PermissionSetMemberClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetMemberClusterType) UnmarshalJSON(b []byte) error {
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
