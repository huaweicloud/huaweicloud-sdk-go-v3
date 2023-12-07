package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSecurityPermissionSetResponse Response Object
type CreateSecurityPermissionSetResponse struct {

	// 编号
	Id *string `json:"id,omitempty"`

	// 父权限集id
	ParentId *string `json:"parent_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 权限集类型, COMMON, MRS_MANAGED
	Type *CreateSecurityPermissionSetResponseType `json:"type,omitempty"`

	// 纳管角色所在集群id（仅纳管类权限集需要）
	ManagedClusterId *string `json:"managed_cluster_id,omitempty"`

	// 纳管角色所在集群名称（仅纳管类权限集需要）
	ManagedClusterName *string `json:"managed_cluster_name,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 管理员id
	ManagerId *string `json:"manager_id,omitempty"`

	// 管理员名称
	ManagerName *string `json:"manager_name,omitempty"`

	// 管理员类型
	ManagerType *string `json:"manager_type,omitempty"`

	// 数据源类型
	DatasourceType *string `json:"datasource_type,omitempty"`

	// 同步状态,UNKNOWN,NOT_SYNC,SYNCING,SYNC_SUCCESS,SYNC_FAIL
	SyncStatus *CreateSecurityPermissionSetResponseSyncStatus `json:"sync_status,omitempty"`

	// 同步信息
	SyncMsg *string `json:"sync_msg,omitempty"`

	// 同步时间
	SyncTime *int64 `json:"sync_time,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 更新者
	UpdateUser     *string `json:"update_user,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSecurityPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityPermissionSetResponse", string(data)}, " ")
}

type CreateSecurityPermissionSetResponseType struct {
	value string
}

type CreateSecurityPermissionSetResponseTypeEnum struct {
	COMMON      CreateSecurityPermissionSetResponseType
	MRS_MANAGED CreateSecurityPermissionSetResponseType
}

func GetCreateSecurityPermissionSetResponseTypeEnum() CreateSecurityPermissionSetResponseTypeEnum {
	return CreateSecurityPermissionSetResponseTypeEnum{
		COMMON: CreateSecurityPermissionSetResponseType{
			value: "COMMON",
		},
		MRS_MANAGED: CreateSecurityPermissionSetResponseType{
			value: "MRS_MANAGED",
		},
	}
}

func (c CreateSecurityPermissionSetResponseType) Value() string {
	return c.value
}

func (c CreateSecurityPermissionSetResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityPermissionSetResponseType) UnmarshalJSON(b []byte) error {
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

type CreateSecurityPermissionSetResponseSyncStatus struct {
	value string
}

type CreateSecurityPermissionSetResponseSyncStatusEnum struct {
	UNKNOWN      CreateSecurityPermissionSetResponseSyncStatus
	NOT_SYNC     CreateSecurityPermissionSetResponseSyncStatus
	SYNCING      CreateSecurityPermissionSetResponseSyncStatus
	SYNC_SUCCESS CreateSecurityPermissionSetResponseSyncStatus
	SYNC_FAIL    CreateSecurityPermissionSetResponseSyncStatus
}

func GetCreateSecurityPermissionSetResponseSyncStatusEnum() CreateSecurityPermissionSetResponseSyncStatusEnum {
	return CreateSecurityPermissionSetResponseSyncStatusEnum{
		UNKNOWN: CreateSecurityPermissionSetResponseSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: CreateSecurityPermissionSetResponseSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: CreateSecurityPermissionSetResponseSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: CreateSecurityPermissionSetResponseSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: CreateSecurityPermissionSetResponseSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c CreateSecurityPermissionSetResponseSyncStatus) Value() string {
	return c.value
}

func (c CreateSecurityPermissionSetResponseSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityPermissionSetResponseSyncStatus) UnmarshalJSON(b []byte) error {
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
