package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSecurityPermissionSetPermissionResponse Response Object
type CreateSecurityPermissionSetPermissionResponse struct {

	// id
	Id *string `json:"id,omitempty"`

	// 权限集id
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 权限类型, DENY, ALLOW
	PermissionType *CreateSecurityPermissionSetPermissionResponsePermissionType `json:"permission_type,omitempty"`

	// 权限操作列表
	PermissionAction *string `json:"permission_action,omitempty"`

	// 权限操作列表
	PermissionActions *[]CreateSecurityPermissionSetPermissionResponsePermissionActions `json:"permission_actions,omitempty"`

	// 权限操作编码, 位图
	PermissionActionCode *int32 `json:"permission_action_code,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据源类型, HIVE
	DatasourceType *CreateSecurityPermissionSetPermissionResponseDatasourceType `json:"datasource_type,omitempty"`

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
	SyncStatus *CreateSecurityPermissionSetPermissionResponseSyncStatus `json:"sync_status,omitempty"`

	// 同步信息
	SyncMsg *string `json:"sync_msg,omitempty"`

	// url路径名称。
	Url            *string `json:"url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSecurityPermissionSetPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPermissionSetPermissionResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityPermissionSetPermissionResponse", string(data)}, " ")
}

type CreateSecurityPermissionSetPermissionResponsePermissionType struct {
	value string
}

type CreateSecurityPermissionSetPermissionResponsePermissionTypeEnum struct {
	DENY  CreateSecurityPermissionSetPermissionResponsePermissionType
	ALLOW CreateSecurityPermissionSetPermissionResponsePermissionType
}

func GetCreateSecurityPermissionSetPermissionResponsePermissionTypeEnum() CreateSecurityPermissionSetPermissionResponsePermissionTypeEnum {
	return CreateSecurityPermissionSetPermissionResponsePermissionTypeEnum{
		DENY: CreateSecurityPermissionSetPermissionResponsePermissionType{
			value: "DENY",
		},
		ALLOW: CreateSecurityPermissionSetPermissionResponsePermissionType{
			value: "ALLOW",
		},
	}
}

func (c CreateSecurityPermissionSetPermissionResponsePermissionType) Value() string {
	return c.value
}

func (c CreateSecurityPermissionSetPermissionResponsePermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityPermissionSetPermissionResponsePermissionType) UnmarshalJSON(b []byte) error {
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

type CreateSecurityPermissionSetPermissionResponsePermissionActions struct {
	value string
}

type CreateSecurityPermissionSetPermissionResponsePermissionActionsEnum struct {
	ALL    CreateSecurityPermissionSetPermissionResponsePermissionActions
	SELECT CreateSecurityPermissionSetPermissionResponsePermissionActions
	UPDATE CreateSecurityPermissionSetPermissionResponsePermissionActions
	CREATE CreateSecurityPermissionSetPermissionResponsePermissionActions
	DROP   CreateSecurityPermissionSetPermissionResponsePermissionActions
	ALTER  CreateSecurityPermissionSetPermissionResponsePermissionActions
	INDEX  CreateSecurityPermissionSetPermissionResponsePermissionActions
	LOCK   CreateSecurityPermissionSetPermissionResponsePermissionActions
	READ   CreateSecurityPermissionSetPermissionResponsePermissionActions
	WRITE  CreateSecurityPermissionSetPermissionResponsePermissionActions
}

func GetCreateSecurityPermissionSetPermissionResponsePermissionActionsEnum() CreateSecurityPermissionSetPermissionResponsePermissionActionsEnum {
	return CreateSecurityPermissionSetPermissionResponsePermissionActionsEnum{
		ALL: CreateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "ALL",
		},
		SELECT: CreateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "SELECT",
		},
		UPDATE: CreateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "UPDATE",
		},
		CREATE: CreateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "CREATE",
		},
		DROP: CreateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "DROP",
		},
		ALTER: CreateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "ALTER",
		},
		INDEX: CreateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "INDEX",
		},
		LOCK: CreateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "LOCK",
		},
		READ: CreateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "READ",
		},
		WRITE: CreateSecurityPermissionSetPermissionResponsePermissionActions{
			value: "WRITE",
		},
	}
}

func (c CreateSecurityPermissionSetPermissionResponsePermissionActions) Value() string {
	return c.value
}

func (c CreateSecurityPermissionSetPermissionResponsePermissionActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityPermissionSetPermissionResponsePermissionActions) UnmarshalJSON(b []byte) error {
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

type CreateSecurityPermissionSetPermissionResponseDatasourceType struct {
	value string
}

type CreateSecurityPermissionSetPermissionResponseDatasourceTypeEnum struct {
	HIVE CreateSecurityPermissionSetPermissionResponseDatasourceType
}

func GetCreateSecurityPermissionSetPermissionResponseDatasourceTypeEnum() CreateSecurityPermissionSetPermissionResponseDatasourceTypeEnum {
	return CreateSecurityPermissionSetPermissionResponseDatasourceTypeEnum{
		HIVE: CreateSecurityPermissionSetPermissionResponseDatasourceType{
			value: "HIVE",
		},
	}
}

func (c CreateSecurityPermissionSetPermissionResponseDatasourceType) Value() string {
	return c.value
}

func (c CreateSecurityPermissionSetPermissionResponseDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityPermissionSetPermissionResponseDatasourceType) UnmarshalJSON(b []byte) error {
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

type CreateSecurityPermissionSetPermissionResponseSyncStatus struct {
	value string
}

type CreateSecurityPermissionSetPermissionResponseSyncStatusEnum struct {
	UNKNOWN      CreateSecurityPermissionSetPermissionResponseSyncStatus
	NOT_SYNC     CreateSecurityPermissionSetPermissionResponseSyncStatus
	SYNCING      CreateSecurityPermissionSetPermissionResponseSyncStatus
	SYNC_SUCCESS CreateSecurityPermissionSetPermissionResponseSyncStatus
	SYNC_FAIL    CreateSecurityPermissionSetPermissionResponseSyncStatus
}

func GetCreateSecurityPermissionSetPermissionResponseSyncStatusEnum() CreateSecurityPermissionSetPermissionResponseSyncStatusEnum {
	return CreateSecurityPermissionSetPermissionResponseSyncStatusEnum{
		UNKNOWN: CreateSecurityPermissionSetPermissionResponseSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: CreateSecurityPermissionSetPermissionResponseSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: CreateSecurityPermissionSetPermissionResponseSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: CreateSecurityPermissionSetPermissionResponseSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: CreateSecurityPermissionSetPermissionResponseSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c CreateSecurityPermissionSetPermissionResponseSyncStatus) Value() string {
	return c.value
}

func (c CreateSecurityPermissionSetPermissionResponseSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityPermissionSetPermissionResponseSyncStatus) UnmarshalJSON(b []byte) error {
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
