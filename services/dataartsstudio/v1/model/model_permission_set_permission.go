package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionSetPermission struct {

	// id
	Id *string `json:"id,omitempty"`

	// 权限集id
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 权限类型, DENY, ALLOW
	PermissionType *PermissionSetPermissionPermissionType `json:"permission_type,omitempty"`

	// 权限操作列表
	PermissionAction *string `json:"permission_action,omitempty"`

	// 权限操作列表
	PermissionActions *[]PermissionSetPermissionPermissionActions `json:"permission_actions,omitempty"`

	// 权限操作编码, 位图
	PermissionActionCode *int32 `json:"permission_action_code,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据源类型, HIVE
	DatasourceType *PermissionSetPermissionDatasourceType `json:"datasource_type,omitempty"`

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
	SyncStatus *PermissionSetPermissionSyncStatus `json:"sync_status,omitempty"`

	// 同步信息
	SyncMsg *string `json:"sync_msg,omitempty"`

	// url路径名称。
	Url *string `json:"url,omitempty"`
}

func (o PermissionSetPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSetPermission struct{}"
	}

	return strings.Join([]string{"PermissionSetPermission", string(data)}, " ")
}

type PermissionSetPermissionPermissionType struct {
	value string
}

type PermissionSetPermissionPermissionTypeEnum struct {
	DENY  PermissionSetPermissionPermissionType
	ALLOW PermissionSetPermissionPermissionType
}

func GetPermissionSetPermissionPermissionTypeEnum() PermissionSetPermissionPermissionTypeEnum {
	return PermissionSetPermissionPermissionTypeEnum{
		DENY: PermissionSetPermissionPermissionType{
			value: "DENY",
		},
		ALLOW: PermissionSetPermissionPermissionType{
			value: "ALLOW",
		},
	}
}

func (c PermissionSetPermissionPermissionType) Value() string {
	return c.value
}

func (c PermissionSetPermissionPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetPermissionPermissionType) UnmarshalJSON(b []byte) error {
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

type PermissionSetPermissionPermissionActions struct {
	value string
}

type PermissionSetPermissionPermissionActionsEnum struct {
	ALL    PermissionSetPermissionPermissionActions
	SELECT PermissionSetPermissionPermissionActions
	UPDATE PermissionSetPermissionPermissionActions
	CREATE PermissionSetPermissionPermissionActions
	DROP   PermissionSetPermissionPermissionActions
	ALTER  PermissionSetPermissionPermissionActions
	INDEX  PermissionSetPermissionPermissionActions
	LOCK   PermissionSetPermissionPermissionActions
	READ   PermissionSetPermissionPermissionActions
	WRITE  PermissionSetPermissionPermissionActions
}

func GetPermissionSetPermissionPermissionActionsEnum() PermissionSetPermissionPermissionActionsEnum {
	return PermissionSetPermissionPermissionActionsEnum{
		ALL: PermissionSetPermissionPermissionActions{
			value: "ALL",
		},
		SELECT: PermissionSetPermissionPermissionActions{
			value: "SELECT",
		},
		UPDATE: PermissionSetPermissionPermissionActions{
			value: "UPDATE",
		},
		CREATE: PermissionSetPermissionPermissionActions{
			value: "CREATE",
		},
		DROP: PermissionSetPermissionPermissionActions{
			value: "DROP",
		},
		ALTER: PermissionSetPermissionPermissionActions{
			value: "ALTER",
		},
		INDEX: PermissionSetPermissionPermissionActions{
			value: "INDEX",
		},
		LOCK: PermissionSetPermissionPermissionActions{
			value: "LOCK",
		},
		READ: PermissionSetPermissionPermissionActions{
			value: "READ",
		},
		WRITE: PermissionSetPermissionPermissionActions{
			value: "WRITE",
		},
	}
}

func (c PermissionSetPermissionPermissionActions) Value() string {
	return c.value
}

func (c PermissionSetPermissionPermissionActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetPermissionPermissionActions) UnmarshalJSON(b []byte) error {
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

type PermissionSetPermissionDatasourceType struct {
	value string
}

type PermissionSetPermissionDatasourceTypeEnum struct {
	HIVE PermissionSetPermissionDatasourceType
}

func GetPermissionSetPermissionDatasourceTypeEnum() PermissionSetPermissionDatasourceTypeEnum {
	return PermissionSetPermissionDatasourceTypeEnum{
		HIVE: PermissionSetPermissionDatasourceType{
			value: "HIVE",
		},
	}
}

func (c PermissionSetPermissionDatasourceType) Value() string {
	return c.value
}

func (c PermissionSetPermissionDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetPermissionDatasourceType) UnmarshalJSON(b []byte) error {
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

type PermissionSetPermissionSyncStatus struct {
	value string
}

type PermissionSetPermissionSyncStatusEnum struct {
	UNKNOWN      PermissionSetPermissionSyncStatus
	NOT_SYNC     PermissionSetPermissionSyncStatus
	SYNCING      PermissionSetPermissionSyncStatus
	SYNC_SUCCESS PermissionSetPermissionSyncStatus
	SYNC_FAIL    PermissionSetPermissionSyncStatus
}

func GetPermissionSetPermissionSyncStatusEnum() PermissionSetPermissionSyncStatusEnum {
	return PermissionSetPermissionSyncStatusEnum{
		UNKNOWN: PermissionSetPermissionSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: PermissionSetPermissionSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: PermissionSetPermissionSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: PermissionSetPermissionSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: PermissionSetPermissionSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c PermissionSetPermissionSyncStatus) Value() string {
	return c.value
}

func (c PermissionSetPermissionSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetPermissionSyncStatus) UnmarshalJSON(b []byte) error {
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
