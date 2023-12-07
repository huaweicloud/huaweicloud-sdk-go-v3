package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionSet struct {

	// 编号
	Id *string `json:"id,omitempty"`

	// 父权限集id
	ParentId *string `json:"parent_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 权限集类型, COMMON, MRS_MANAGED
	Type *PermissionSetType `json:"type,omitempty"`

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
	SyncStatus *PermissionSetSyncStatus `json:"sync_status,omitempty"`

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
	UpdateUser *string `json:"update_user,omitempty"`
}

func (o PermissionSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSet struct{}"
	}

	return strings.Join([]string{"PermissionSet", string(data)}, " ")
}

type PermissionSetType struct {
	value string
}

type PermissionSetTypeEnum struct {
	COMMON      PermissionSetType
	MRS_MANAGED PermissionSetType
}

func GetPermissionSetTypeEnum() PermissionSetTypeEnum {
	return PermissionSetTypeEnum{
		COMMON: PermissionSetType{
			value: "COMMON",
		},
		MRS_MANAGED: PermissionSetType{
			value: "MRS_MANAGED",
		},
	}
}

func (c PermissionSetType) Value() string {
	return c.value
}

func (c PermissionSetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetType) UnmarshalJSON(b []byte) error {
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

type PermissionSetSyncStatus struct {
	value string
}

type PermissionSetSyncStatusEnum struct {
	UNKNOWN      PermissionSetSyncStatus
	NOT_SYNC     PermissionSetSyncStatus
	SYNCING      PermissionSetSyncStatus
	SYNC_SUCCESS PermissionSetSyncStatus
	SYNC_FAIL    PermissionSetSyncStatus
}

func GetPermissionSetSyncStatusEnum() PermissionSetSyncStatusEnum {
	return PermissionSetSyncStatusEnum{
		UNKNOWN: PermissionSetSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: PermissionSetSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: PermissionSetSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: PermissionSetSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: PermissionSetSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c PermissionSetSyncStatus) Value() string {
	return c.value
}

func (c PermissionSetSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetSyncStatus) UnmarshalJSON(b []byte) error {
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
