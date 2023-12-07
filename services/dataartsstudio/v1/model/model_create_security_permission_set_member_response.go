package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSecurityPermissionSetMemberResponse Response Object
type CreateSecurityPermissionSetMemberResponse struct {

	// id
	Id *string `json:"id,omitempty"`

	// 权限集id
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 成员类型, 用户/用户组/工作空间角色(废弃)/集群角色, USER, USER_GROUP, WORKSPACE_ROLE, CLUSTER_ROLE
	MemberType *CreateSecurityPermissionSetMemberResponseMemberType `json:"member_type,omitempty"`

	// 成员id
	MemberId *string `json:"member_id,omitempty"`

	// 成员名称
	MemberName *string `json:"member_name,omitempty"`

	// 成员状态, NORMAL, UNFINISHED
	MemberStatus *CreateSecurityPermissionSetMemberResponseMemberStatus `json:"member_status,omitempty"`

	// 工作空间(仅工作空间角色需要)
	Workspace *string `json:"workspace,omitempty"`

	// 集群类型(仅集群角色需要), MRS, DWS, DLI
	ClusterType *CreateSecurityPermissionSetMemberResponseClusterType `json:"cluster_type,omitempty"`

	// 集群id(仅集群角色需要)
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称(仅集群角色需要)
	ClusterName *string `json:"cluster_name,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 到期时间
	Deadline       *int64 `json:"deadline,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateSecurityPermissionSetMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPermissionSetMemberResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityPermissionSetMemberResponse", string(data)}, " ")
}

type CreateSecurityPermissionSetMemberResponseMemberType struct {
	value string
}

type CreateSecurityPermissionSetMemberResponseMemberTypeEnum struct {
	USER           CreateSecurityPermissionSetMemberResponseMemberType
	USER_GROUP     CreateSecurityPermissionSetMemberResponseMemberType
	WORKSPACE_ROLE CreateSecurityPermissionSetMemberResponseMemberType
	CLUSTER_ROLE   CreateSecurityPermissionSetMemberResponseMemberType
}

func GetCreateSecurityPermissionSetMemberResponseMemberTypeEnum() CreateSecurityPermissionSetMemberResponseMemberTypeEnum {
	return CreateSecurityPermissionSetMemberResponseMemberTypeEnum{
		USER: CreateSecurityPermissionSetMemberResponseMemberType{
			value: "USER",
		},
		USER_GROUP: CreateSecurityPermissionSetMemberResponseMemberType{
			value: "USER_GROUP",
		},
		WORKSPACE_ROLE: CreateSecurityPermissionSetMemberResponseMemberType{
			value: "WORKSPACE_ROLE",
		},
		CLUSTER_ROLE: CreateSecurityPermissionSetMemberResponseMemberType{
			value: "CLUSTER_ROLE",
		},
	}
}

func (c CreateSecurityPermissionSetMemberResponseMemberType) Value() string {
	return c.value
}

func (c CreateSecurityPermissionSetMemberResponseMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityPermissionSetMemberResponseMemberType) UnmarshalJSON(b []byte) error {
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

type CreateSecurityPermissionSetMemberResponseMemberStatus struct {
	value string
}

type CreateSecurityPermissionSetMemberResponseMemberStatusEnum struct {
	NORMAL     CreateSecurityPermissionSetMemberResponseMemberStatus
	UNFINISHED CreateSecurityPermissionSetMemberResponseMemberStatus
}

func GetCreateSecurityPermissionSetMemberResponseMemberStatusEnum() CreateSecurityPermissionSetMemberResponseMemberStatusEnum {
	return CreateSecurityPermissionSetMemberResponseMemberStatusEnum{
		NORMAL: CreateSecurityPermissionSetMemberResponseMemberStatus{
			value: "NORMAL",
		},
		UNFINISHED: CreateSecurityPermissionSetMemberResponseMemberStatus{
			value: "UNFINISHED",
		},
	}
}

func (c CreateSecurityPermissionSetMemberResponseMemberStatus) Value() string {
	return c.value
}

func (c CreateSecurityPermissionSetMemberResponseMemberStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityPermissionSetMemberResponseMemberStatus) UnmarshalJSON(b []byte) error {
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

type CreateSecurityPermissionSetMemberResponseClusterType struct {
	value string
}

type CreateSecurityPermissionSetMemberResponseClusterTypeEnum struct {
	MRS CreateSecurityPermissionSetMemberResponseClusterType
	DWS CreateSecurityPermissionSetMemberResponseClusterType
	DLI CreateSecurityPermissionSetMemberResponseClusterType
}

func GetCreateSecurityPermissionSetMemberResponseClusterTypeEnum() CreateSecurityPermissionSetMemberResponseClusterTypeEnum {
	return CreateSecurityPermissionSetMemberResponseClusterTypeEnum{
		MRS: CreateSecurityPermissionSetMemberResponseClusterType{
			value: "MRS",
		},
		DWS: CreateSecurityPermissionSetMemberResponseClusterType{
			value: "DWS",
		},
		DLI: CreateSecurityPermissionSetMemberResponseClusterType{
			value: "DLI",
		},
	}
}

func (c CreateSecurityPermissionSetMemberResponseClusterType) Value() string {
	return c.value
}

func (c CreateSecurityPermissionSetMemberResponseClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityPermissionSetMemberResponseClusterType) UnmarshalJSON(b []byte) error {
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
