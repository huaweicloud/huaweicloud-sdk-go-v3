package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSecurityPermissionSetResponse Response Object
type UpdateSecurityPermissionSetResponse struct {

	// 编号
	Id *string `json:"id,omitempty"`

	// 父权限集id
	ParentId *string `json:"parent_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 权限集类型, COMMON, MRS_MANAGED
	Type *UpdateSecurityPermissionSetResponseType `json:"type,omitempty"`

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
	SyncStatus *UpdateSecurityPermissionSetResponseSyncStatus `json:"sync_status,omitempty"`

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

func (o UpdateSecurityPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPermissionSetResponse", string(data)}, " ")
}

type UpdateSecurityPermissionSetResponseType struct {
	value string
}

type UpdateSecurityPermissionSetResponseTypeEnum struct {
	COMMON      UpdateSecurityPermissionSetResponseType
	MRS_MANAGED UpdateSecurityPermissionSetResponseType
}

func GetUpdateSecurityPermissionSetResponseTypeEnum() UpdateSecurityPermissionSetResponseTypeEnum {
	return UpdateSecurityPermissionSetResponseTypeEnum{
		COMMON: UpdateSecurityPermissionSetResponseType{
			value: "COMMON",
		},
		MRS_MANAGED: UpdateSecurityPermissionSetResponseType{
			value: "MRS_MANAGED",
		},
	}
}

func (c UpdateSecurityPermissionSetResponseType) Value() string {
	return c.value
}

func (c UpdateSecurityPermissionSetResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecurityPermissionSetResponseType) UnmarshalJSON(b []byte) error {
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

type UpdateSecurityPermissionSetResponseSyncStatus struct {
	value string
}

type UpdateSecurityPermissionSetResponseSyncStatusEnum struct {
	UNKNOWN      UpdateSecurityPermissionSetResponseSyncStatus
	NOT_SYNC     UpdateSecurityPermissionSetResponseSyncStatus
	SYNCING      UpdateSecurityPermissionSetResponseSyncStatus
	SYNC_SUCCESS UpdateSecurityPermissionSetResponseSyncStatus
	SYNC_FAIL    UpdateSecurityPermissionSetResponseSyncStatus
}

func GetUpdateSecurityPermissionSetResponseSyncStatusEnum() UpdateSecurityPermissionSetResponseSyncStatusEnum {
	return UpdateSecurityPermissionSetResponseSyncStatusEnum{
		UNKNOWN: UpdateSecurityPermissionSetResponseSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: UpdateSecurityPermissionSetResponseSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: UpdateSecurityPermissionSetResponseSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: UpdateSecurityPermissionSetResponseSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: UpdateSecurityPermissionSetResponseSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c UpdateSecurityPermissionSetResponseSyncStatus) Value() string {
	return c.value
}

func (c UpdateSecurityPermissionSetResponseSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecurityPermissionSetResponseSyncStatus) UnmarshalJSON(b []byte) error {
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
