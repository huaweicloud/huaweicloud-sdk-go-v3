package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MemberPermission 用户权限详情
type MemberPermission struct {

	// 权限集ID
	PermissionSetId string `json:"permission_set_id"`

	// 权限来源：1、权限集名称。2、权限审批
	PermissionSource string `json:"permission_source"`

	// 权限类别,ALL,SELECT,UPDATE,CREATE,DROP,ALTER,INDEX,LOCK,READ,WRITE
	PermissionActions MemberPermissionPermissionActions `json:"permission_actions"`

	// Hive数据源，指定url权限的策略信息
	Url *string `json:"url,omitempty"`

	// 数据源类型
	DatasourceType string `json:"datasource_type"`

	// 集群名
	ClusterName string `json:"cluster_name"`

	// 数据库名
	DatabaseName string `json:"database_name"`

	// schema名
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名
	TableName string `json:"table_name"`

	// 列名
	ColumnName *string `json:"column_name,omitempty"`
}

func (o MemberPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberPermission struct{}"
	}

	return strings.Join([]string{"MemberPermission", string(data)}, " ")
}

type MemberPermissionPermissionActions struct {
	value string
}

type MemberPermissionPermissionActionsEnum struct {
	ALL    MemberPermissionPermissionActions
	SELECT MemberPermissionPermissionActions
	UPDATE MemberPermissionPermissionActions
	CREATE MemberPermissionPermissionActions
	DROP   MemberPermissionPermissionActions
	ALTER  MemberPermissionPermissionActions
	INDEX  MemberPermissionPermissionActions
	LOCK   MemberPermissionPermissionActions
	READ   MemberPermissionPermissionActions
	WRITE  MemberPermissionPermissionActions
}

func GetMemberPermissionPermissionActionsEnum() MemberPermissionPermissionActionsEnum {
	return MemberPermissionPermissionActionsEnum{
		ALL: MemberPermissionPermissionActions{
			value: "ALL",
		},
		SELECT: MemberPermissionPermissionActions{
			value: "SELECT",
		},
		UPDATE: MemberPermissionPermissionActions{
			value: "UPDATE",
		},
		CREATE: MemberPermissionPermissionActions{
			value: "CREATE",
		},
		DROP: MemberPermissionPermissionActions{
			value: "DROP",
		},
		ALTER: MemberPermissionPermissionActions{
			value: "ALTER",
		},
		INDEX: MemberPermissionPermissionActions{
			value: "INDEX",
		},
		LOCK: MemberPermissionPermissionActions{
			value: "LOCK",
		},
		READ: MemberPermissionPermissionActions{
			value: "READ",
		},
		WRITE: MemberPermissionPermissionActions{
			value: "WRITE",
		},
	}
}

func (c MemberPermissionPermissionActions) Value() string {
	return c.value
}

func (c MemberPermissionPermissionActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberPermissionPermissionActions) UnmarshalJSON(b []byte) error {
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
