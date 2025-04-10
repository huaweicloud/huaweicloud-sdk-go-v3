package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccountPermission 用户权限详情
type AccountPermission struct {

	// 集群id
	ClusterId string `json:"cluster_id"`

	// 集群名称
	ClusterName string `json:"cluster_name"`

	// 列名
	ColumnName *string `json:"column_name,omitempty"`

	// 数据库名
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据源类型,HIVE
	DatasourceType AccountPermissionDatasourceType `json:"datasource_type"`

	// 到期信息
	ExpireMsg *string `json:"expire_msg,omitempty"`

	// 权限状态,REVOKE_FAILED,TO_BE_REVOKE,INACTIVE,PERMANENTLY_ACTIVE,ACTIVE,EXPIRE_SOON
	ExpireStatus AccountPermissionExpireStatus `json:"expire_status"`

	// 到期时间
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 权限id
	Id string `json:"id"`

	// 实例id
	InstanceId string `json:"instance_id"`

	// 成员id
	MemberId *string `json:"member_id,omitempty"`

	// 成员名称
	MemberName *string `json:"member_name,omitempty"`

	// 权限类别,ALL,SELECT,UPDATE,CREATE,DROP,ALTER,INDEX,LOCK,READ,WRITE
	PermissionAction *AccountPermissionPermissionAction `json:"permission_action,omitempty"`

	// 权限位图
	PermissionActionCode *int64 `json:"permission_action_code,omitempty"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 行级权限表达式
	RowLevelSecurity *string `json:"row_level_security,omitempty"`

	// 行级权限描述
	RowLevelSecurityDesc *string `json:"row_level_security_desc,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`
}

func (o AccountPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountPermission struct{}"
	}

	return strings.Join([]string{"AccountPermission", string(data)}, " ")
}

type AccountPermissionDatasourceType struct {
	value string
}

type AccountPermissionDatasourceTypeEnum struct {
	HIVE AccountPermissionDatasourceType
}

func GetAccountPermissionDatasourceTypeEnum() AccountPermissionDatasourceTypeEnum {
	return AccountPermissionDatasourceTypeEnum{
		HIVE: AccountPermissionDatasourceType{
			value: "HIVE",
		},
	}
}

func (c AccountPermissionDatasourceType) Value() string {
	return c.value
}

func (c AccountPermissionDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountPermissionDatasourceType) UnmarshalJSON(b []byte) error {
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

type AccountPermissionExpireStatus struct {
	value string
}

type AccountPermissionExpireStatusEnum struct {
	REVOKE_FAILED      AccountPermissionExpireStatus
	TO_BE_REVOKE       AccountPermissionExpireStatus
	INACTIVE           AccountPermissionExpireStatus
	PERMANENTLY_ACTIVE AccountPermissionExpireStatus
	ACTIVE             AccountPermissionExpireStatus
	EXPIRE_SOON        AccountPermissionExpireStatus
}

func GetAccountPermissionExpireStatusEnum() AccountPermissionExpireStatusEnum {
	return AccountPermissionExpireStatusEnum{
		REVOKE_FAILED: AccountPermissionExpireStatus{
			value: "REVOKE_FAILED",
		},
		TO_BE_REVOKE: AccountPermissionExpireStatus{
			value: "TO_BE_REVOKE",
		},
		INACTIVE: AccountPermissionExpireStatus{
			value: "INACTIVE",
		},
		PERMANENTLY_ACTIVE: AccountPermissionExpireStatus{
			value: "PERMANENTLY_ACTIVE",
		},
		ACTIVE: AccountPermissionExpireStatus{
			value: "ACTIVE",
		},
		EXPIRE_SOON: AccountPermissionExpireStatus{
			value: "EXPIRE_SOON",
		},
	}
}

func (c AccountPermissionExpireStatus) Value() string {
	return c.value
}

func (c AccountPermissionExpireStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountPermissionExpireStatus) UnmarshalJSON(b []byte) error {
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

type AccountPermissionPermissionAction struct {
	value string
}

type AccountPermissionPermissionActionEnum struct {
	ALL    AccountPermissionPermissionAction
	SELECT AccountPermissionPermissionAction
	UPDATE AccountPermissionPermissionAction
	CREATE AccountPermissionPermissionAction
	DROP   AccountPermissionPermissionAction
	ALTER  AccountPermissionPermissionAction
	INDEX  AccountPermissionPermissionAction
	LOCK   AccountPermissionPermissionAction
	READ   AccountPermissionPermissionAction
	WRITE  AccountPermissionPermissionAction
}

func GetAccountPermissionPermissionActionEnum() AccountPermissionPermissionActionEnum {
	return AccountPermissionPermissionActionEnum{
		ALL: AccountPermissionPermissionAction{
			value: "ALL",
		},
		SELECT: AccountPermissionPermissionAction{
			value: "SELECT",
		},
		UPDATE: AccountPermissionPermissionAction{
			value: "UPDATE",
		},
		CREATE: AccountPermissionPermissionAction{
			value: "CREATE",
		},
		DROP: AccountPermissionPermissionAction{
			value: "DROP",
		},
		ALTER: AccountPermissionPermissionAction{
			value: "ALTER",
		},
		INDEX: AccountPermissionPermissionAction{
			value: "INDEX",
		},
		LOCK: AccountPermissionPermissionAction{
			value: "LOCK",
		},
		READ: AccountPermissionPermissionAction{
			value: "READ",
		},
		WRITE: AccountPermissionPermissionAction{
			value: "WRITE",
		},
	}
}

func (c AccountPermissionPermissionAction) Value() string {
	return c.value
}

func (c AccountPermissionPermissionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountPermissionPermissionAction) UnmarshalJSON(b []byte) error {
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
