package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityPermissionSetPermissionsRequest Request Object
type ListSecurityPermissionSetPermissionsRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// workspace 信息
	Workspace string `json:"workspace"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 权限类型,DENY,ALLOW
	PermissionType *ListSecurityPermissionSetPermissionsRequestPermissionType `json:"permission_type,omitempty"`

	// 权限操作,ALL,SELECT,UPDATE,CREATE,DROP,ALTER,INDEX,LOCK,READ,WRITE
	PermissionAction *ListSecurityPermissionSetPermissionsRequestPermissionAction `json:"permission_action,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据源类型,HIVE
	DatasourceType *ListSecurityPermissionSetPermissionsRequestDatasourceType `json:"datasource_type,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 同步状态,UNKNOWN,NOT_SYNC,SYNCING,SYNC_SUCCESS,SYNC_FAIL
	SyncStatus *ListSecurityPermissionSetPermissionsRequestSyncStatus `json:"sync_status,omitempty"`

	// 排序参数, CLUSTER_NAME, DATABASE_NAME
	OrderBy *ListSecurityPermissionSetPermissionsRequestOrderBy `json:"order_by,omitempty"`

	// 是否升序（仅指定排序参数时有效）
	OrderByAsc *bool `json:"order_by_asc,omitempty"`
}

func (o ListSecurityPermissionSetPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPermissionSetPermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityPermissionSetPermissionsRequest", string(data)}, " ")
}

type ListSecurityPermissionSetPermissionsRequestPermissionType struct {
	value string
}

type ListSecurityPermissionSetPermissionsRequestPermissionTypeEnum struct {
	DENY  ListSecurityPermissionSetPermissionsRequestPermissionType
	ALLOW ListSecurityPermissionSetPermissionsRequestPermissionType
}

func GetListSecurityPermissionSetPermissionsRequestPermissionTypeEnum() ListSecurityPermissionSetPermissionsRequestPermissionTypeEnum {
	return ListSecurityPermissionSetPermissionsRequestPermissionTypeEnum{
		DENY: ListSecurityPermissionSetPermissionsRequestPermissionType{
			value: "DENY",
		},
		ALLOW: ListSecurityPermissionSetPermissionsRequestPermissionType{
			value: "ALLOW",
		},
	}
}

func (c ListSecurityPermissionSetPermissionsRequestPermissionType) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetPermissionsRequestPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetPermissionsRequestPermissionType) UnmarshalJSON(b []byte) error {
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

type ListSecurityPermissionSetPermissionsRequestPermissionAction struct {
	value string
}

type ListSecurityPermissionSetPermissionsRequestPermissionActionEnum struct {
	ALL    ListSecurityPermissionSetPermissionsRequestPermissionAction
	SELECT ListSecurityPermissionSetPermissionsRequestPermissionAction
	UPDATE ListSecurityPermissionSetPermissionsRequestPermissionAction
	CREATE ListSecurityPermissionSetPermissionsRequestPermissionAction
	DROP   ListSecurityPermissionSetPermissionsRequestPermissionAction
	ALTER  ListSecurityPermissionSetPermissionsRequestPermissionAction
	INDEX  ListSecurityPermissionSetPermissionsRequestPermissionAction
	LOCK   ListSecurityPermissionSetPermissionsRequestPermissionAction
	READ   ListSecurityPermissionSetPermissionsRequestPermissionAction
	WRITE  ListSecurityPermissionSetPermissionsRequestPermissionAction
}

func GetListSecurityPermissionSetPermissionsRequestPermissionActionEnum() ListSecurityPermissionSetPermissionsRequestPermissionActionEnum {
	return ListSecurityPermissionSetPermissionsRequestPermissionActionEnum{
		ALL: ListSecurityPermissionSetPermissionsRequestPermissionAction{
			value: "ALL",
		},
		SELECT: ListSecurityPermissionSetPermissionsRequestPermissionAction{
			value: "SELECT",
		},
		UPDATE: ListSecurityPermissionSetPermissionsRequestPermissionAction{
			value: "UPDATE",
		},
		CREATE: ListSecurityPermissionSetPermissionsRequestPermissionAction{
			value: "CREATE",
		},
		DROP: ListSecurityPermissionSetPermissionsRequestPermissionAction{
			value: "DROP",
		},
		ALTER: ListSecurityPermissionSetPermissionsRequestPermissionAction{
			value: "ALTER",
		},
		INDEX: ListSecurityPermissionSetPermissionsRequestPermissionAction{
			value: "INDEX",
		},
		LOCK: ListSecurityPermissionSetPermissionsRequestPermissionAction{
			value: "LOCK",
		},
		READ: ListSecurityPermissionSetPermissionsRequestPermissionAction{
			value: "READ",
		},
		WRITE: ListSecurityPermissionSetPermissionsRequestPermissionAction{
			value: "WRITE",
		},
	}
}

func (c ListSecurityPermissionSetPermissionsRequestPermissionAction) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetPermissionsRequestPermissionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetPermissionsRequestPermissionAction) UnmarshalJSON(b []byte) error {
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

type ListSecurityPermissionSetPermissionsRequestDatasourceType struct {
	value string
}

type ListSecurityPermissionSetPermissionsRequestDatasourceTypeEnum struct {
	HIVE ListSecurityPermissionSetPermissionsRequestDatasourceType
}

func GetListSecurityPermissionSetPermissionsRequestDatasourceTypeEnum() ListSecurityPermissionSetPermissionsRequestDatasourceTypeEnum {
	return ListSecurityPermissionSetPermissionsRequestDatasourceTypeEnum{
		HIVE: ListSecurityPermissionSetPermissionsRequestDatasourceType{
			value: "HIVE",
		},
	}
}

func (c ListSecurityPermissionSetPermissionsRequestDatasourceType) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetPermissionsRequestDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetPermissionsRequestDatasourceType) UnmarshalJSON(b []byte) error {
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

type ListSecurityPermissionSetPermissionsRequestSyncStatus struct {
	value string
}

type ListSecurityPermissionSetPermissionsRequestSyncStatusEnum struct {
	UNKNOWN      ListSecurityPermissionSetPermissionsRequestSyncStatus
	NOT_SYNC     ListSecurityPermissionSetPermissionsRequestSyncStatus
	SYNCING      ListSecurityPermissionSetPermissionsRequestSyncStatus
	SYNC_SUCCESS ListSecurityPermissionSetPermissionsRequestSyncStatus
	SYNC_FAIL    ListSecurityPermissionSetPermissionsRequestSyncStatus
}

func GetListSecurityPermissionSetPermissionsRequestSyncStatusEnum() ListSecurityPermissionSetPermissionsRequestSyncStatusEnum {
	return ListSecurityPermissionSetPermissionsRequestSyncStatusEnum{
		UNKNOWN: ListSecurityPermissionSetPermissionsRequestSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: ListSecurityPermissionSetPermissionsRequestSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: ListSecurityPermissionSetPermissionsRequestSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: ListSecurityPermissionSetPermissionsRequestSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: ListSecurityPermissionSetPermissionsRequestSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c ListSecurityPermissionSetPermissionsRequestSyncStatus) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetPermissionsRequestSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetPermissionsRequestSyncStatus) UnmarshalJSON(b []byte) error {
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

type ListSecurityPermissionSetPermissionsRequestOrderBy struct {
	value string
}

type ListSecurityPermissionSetPermissionsRequestOrderByEnum struct {
	CLUSTER_NAME  ListSecurityPermissionSetPermissionsRequestOrderBy
	DATABASE_NAME ListSecurityPermissionSetPermissionsRequestOrderBy
}

func GetListSecurityPermissionSetPermissionsRequestOrderByEnum() ListSecurityPermissionSetPermissionsRequestOrderByEnum {
	return ListSecurityPermissionSetPermissionsRequestOrderByEnum{
		CLUSTER_NAME: ListSecurityPermissionSetPermissionsRequestOrderBy{
			value: "CLUSTER_NAME",
		},
		DATABASE_NAME: ListSecurityPermissionSetPermissionsRequestOrderBy{
			value: "DATABASE_NAME",
		},
	}
}

func (c ListSecurityPermissionSetPermissionsRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetPermissionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetPermissionsRequestOrderBy) UnmarshalJSON(b []byte) error {
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
