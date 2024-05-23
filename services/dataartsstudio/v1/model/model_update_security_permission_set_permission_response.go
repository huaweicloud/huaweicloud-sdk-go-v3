package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSecurityPermissionSetPermissionResponse Response Object
type UpdateSecurityPermissionSetPermissionResponse struct {

	// id
	Id *string `json:"id,omitempty"`

	// 权限集id
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 权限类型, DENY, ALLOW
	PermissionType *UpdateSecurityPermissionSetPermissionResponsePermissionType `json:"permission_type,omitempty"`

	// 权限操作列表
	PermissionAction *string `json:"permission_action,omitempty"`

	// 权限操作列表
	PermissionActions *[]UpdateSecurityPermissionSetPermissionResponsePermissionActions `json:"permission_actions,omitempty"`

	// 权限操作编码, 位图
	PermissionActionCode *int32 `json:"permission_action_code,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据源类型, HIVE
	DatasourceType *UpdateSecurityPermissionSetPermissionResponseDatasourceType `json:"datasource_type,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 模式名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 行级策略
	RowLevelSecurity *string `json:"row_level_security,omitempty"`

	// 同步状态,UNKNOWN,NOT_SYNC,SYNCING,SYNC_SUCCESS,SYNC_FAIL
	SyncStatus *UpdateSecurityPermissionSetPermissionResponseSyncStatus `json:"sync_status,omitempty"`

	// 同步信息
	SyncMsg *string `json:"sync_msg,omitempty"`

	// url路径名称。
	Url            *string `json:"url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSecurityPermissionSetPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPermissionSetPermissionResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPermissionSetPermissionResponse", string(data)}, " ")
}

type UpdateSecurityPermissionSetPermissionResponsePermissionType struct {
	value string
}

type UpdateSecurityPermissionSetPermissionResponsePermissionTypeEnum struct {
	DENY  UpdateSecurityPermissionSetPermissionResponsePermissionType
	ALLOW UpdateSecurityPermissionSetPermissionResponsePermissionType
}

func GetUpdateSecurityPermissionSetPermissionResponsePermissionTypeEnum() UpdateSecurityPermissionSetPermissionResponsePermissionTypeEnum {
	return UpdateSecurityPermissionSetPermissionResponsePermissionTypeEnum{
		DENY: UpdateSecurityPermissionSetPermissionResponsePermissionType{
			value: "DENY",
		},
		ALLOW: UpdateSecurityPermissionSetPermissionResponsePermissionType{
			value: "ALLOW",
		},
	}
}

func (c UpdateSecurityPermissionSetPermissionResponsePermissionType) Value() string {
	return c.value
}

func (c UpdateSecurityPermissionSetPermissionResponsePermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecurityPermissionSetPermissionResponsePermissionType) UnmarshalJSON(b []byte) error {
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

type UpdateSecurityPermissionSetPermissionResponsePermissionActions struct {
	value string
}

type UpdateSecurityPermissionSetPermissionResponsePermissionActionsEnum struct {
	ALL    UpdateSecurityPermissionSetPermissionResponsePermissionActions
	SELECT UpdateSecurityPermissionSetPermissionResponsePermissionActions
	UPDATE UpdateSecurityPermissionSetPermissionResponsePermissionActions
	CREATE UpdateSecurityPermissionSetPermissionResponsePermissionActions
	DROP   UpdateSecurityPermissionSetPermissionResponsePermissionActions
	ALTER  UpdateSecurityPermissionSetPermissionResponsePermissionActions
	INDEX  UpdateSecurityPermissionSetPermissionResponsePermissionActions
	LOCK   UpdateSecurityPermissionSetPermissionResponsePermissionActions
	READ   UpdateSecurityPermissionSetPermissionResponsePermissionActions
	WRITE  UpdateSecurityPermissionSetPermissionResponsePermissionActions
}

func GetUpdateSecurityPermissionSetPermissionResponsePermissionActionsEnum() UpdateSecurityPermissionSetPermissionResponsePermissionActionsEnum {
	return UpdateSecurityPermissionSetPermissionResponsePermissionActionsEnum{
		ALL: UpdateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "ALL",
		},
		SELECT: UpdateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "SELECT",
		},
		UPDATE: UpdateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "UPDATE",
		},
		CREATE: UpdateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "CREATE",
		},
		DROP: UpdateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "DROP",
		},
		ALTER: UpdateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "ALTER",
		},
		INDEX: UpdateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "INDEX",
		},
		LOCK: UpdateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "LOCK",
		},
		READ: UpdateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "READ",
		},
		WRITE: UpdateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "WRITE",
		},
	}
}

func (c UpdateSecurityPermissionSetPermissionResponsePermissionActions) Value() string {
	return c.value
}

func (c UpdateSecurityPermissionSetPermissionResponsePermissionActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecurityPermissionSetPermissionResponsePermissionActions) UnmarshalJSON(b []byte) error {
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

type UpdateSecurityPermissionSetPermissionResponseDatasourceType struct {
	value string
}

type UpdateSecurityPermissionSetPermissionResponseDatasourceTypeEnum struct {
	HIVE UpdateSecurityPermissionSetPermissionResponseDatasourceType
}

func GetUpdateSecurityPermissionSetPermissionResponseDatasourceTypeEnum() UpdateSecurityPermissionSetPermissionResponseDatasourceTypeEnum {
	return UpdateSecurityPermissionSetPermissionResponseDatasourceTypeEnum{
		HIVE: UpdateSecurityPermissionSetPermissionResponseDatasourceType{
			value: "HIVE",
		},
	}
}

func (c UpdateSecurityPermissionSetPermissionResponseDatasourceType) Value() string {
	return c.value
}

func (c UpdateSecurityPermissionSetPermissionResponseDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecurityPermissionSetPermissionResponseDatasourceType) UnmarshalJSON(b []byte) error {
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

type UpdateSecurityPermissionSetPermissionResponseSyncStatus struct {
	value string
}

type UpdateSecurityPermissionSetPermissionResponseSyncStatusEnum struct {
	UNKNOWN      UpdateSecurityPermissionSetPermissionResponseSyncStatus
	NOT_SYNC     UpdateSecurityPermissionSetPermissionResponseSyncStatus
	SYNCING      UpdateSecurityPermissionSetPermissionResponseSyncStatus
	SYNC_SUCCESS UpdateSecurityPermissionSetPermissionResponseSyncStatus
	SYNC_FAIL    UpdateSecurityPermissionSetPermissionResponseSyncStatus
}

func GetUpdateSecurityPermissionSetPermissionResponseSyncStatusEnum() UpdateSecurityPermissionSetPermissionResponseSyncStatusEnum {
	return UpdateSecurityPermissionSetPermissionResponseSyncStatusEnum{
		UNKNOWN: UpdateSecurityPermissionSetPermissionResponseSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: UpdateSecurityPermissionSetPermissionResponseSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: UpdateSecurityPermissionSetPermissionResponseSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: UpdateSecurityPermissionSetPermissionResponseSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: UpdateSecurityPermissionSetPermissionResponseSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c UpdateSecurityPermissionSetPermissionResponseSyncStatus) Value() string {
	return c.value
}

func (c UpdateSecurityPermissionSetPermissionResponseSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecurityPermissionSetPermissionResponseSyncStatus) UnmarshalJSON(b []byte) error {
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
