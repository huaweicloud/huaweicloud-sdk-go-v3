package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSecurityPermissionSetResponse Response Object
type ShowSecurityPermissionSetResponse struct {

	// 编号
	Id *string `json:"id,omitempty"`

	// 父权限集id
	ParentId *string `json:"parent_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 权限集类型, COMMON, MRS_MANAGED
	Type *ShowSecurityPermissionSetResponseType `json:"type,omitempty"`

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
	SyncStatus *ShowSecurityPermissionSetResponseSyncStatus `json:"sync_status,omitempty"`

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

func (o ShowSecurityPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityPermissionSetResponse", string(data)}, " ")
}

type ShowSecurityPermissionSetResponseType struct {
	value string
}

type ShowSecurityPermissionSetResponseTypeEnum struct {
	COMMON      ShowSecurityPermissionSetResponseType
	MRS_MANAGED ShowSecurityPermissionSetResponseType
}

func GetShowSecurityPermissionSetResponseTypeEnum() ShowSecurityPermissionSetResponseTypeEnum {
	return ShowSecurityPermissionSetResponseTypeEnum{
		COMMON: ShowSecurityPermissionSetResponseType{
			value: "COMMON",
		},
		MRS_MANAGED: ShowSecurityPermissionSetResponseType{
			value: "MRS_MANAGED",
		},
	}
}

func (c ShowSecurityPermissionSetResponseType) Value() string {
	return c.value
}

func (c ShowSecurityPermissionSetResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityPermissionSetResponseType) UnmarshalJSON(b []byte) error {
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

type ShowSecurityPermissionSetResponseSyncStatus struct {
	value string
}

type ShowSecurityPermissionSetResponseSyncStatusEnum struct {
	UNKNOWN      ShowSecurityPermissionSetResponseSyncStatus
	NOT_SYNC     ShowSecurityPermissionSetResponseSyncStatus
	SYNCING      ShowSecurityPermissionSetResponseSyncStatus
	SYNC_SUCCESS ShowSecurityPermissionSetResponseSyncStatus
	SYNC_FAIL    ShowSecurityPermissionSetResponseSyncStatus
}

func GetShowSecurityPermissionSetResponseSyncStatusEnum() ShowSecurityPermissionSetResponseSyncStatusEnum {
	return ShowSecurityPermissionSetResponseSyncStatusEnum{
		UNKNOWN: ShowSecurityPermissionSetResponseSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: ShowSecurityPermissionSetResponseSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: ShowSecurityPermissionSetResponseSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: ShowSecurityPermissionSetResponseSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: ShowSecurityPermissionSetResponseSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c ShowSecurityPermissionSetResponseSyncStatus) Value() string {
	return c.value
}

func (c ShowSecurityPermissionSetResponseSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityPermissionSetResponseSyncStatus) UnmarshalJSON(b []byte) error {
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
