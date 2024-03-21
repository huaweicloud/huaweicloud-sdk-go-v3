package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PermissionActions HIVE数据源支持权限操作类型：   * `ALL` - 所有执行权限   * `SELECT` - 查询权限   * `UPDATE` - 更新权限   * `CREATE` - 创建权限   * `DROP` - drop操作权限   * `ALTER` - alter操作权限   * `INDEX` - 索引操作权限   * `READ` - 可读权限   * `WRITE` - 可写权限  DLI数据源支持权限操作类型：   * `SELECT` - 查询权限   * `DROP` - drop操作权限   * `ALTER` - alter操作权限   * `INSERT` - 插入数据权限   * `CREATE_TABLE` - 创建表权限  DWS数据源支持权限操作类型：   * `ALL` - 所有执行权限   * `SELECT` - 查询权限   * `UPDATE` - 更新权限   * `DROP` - drop操作权限   * `ALTER` - alter操作权限   * `INSERT` - 插入数据权限   * `CREATE_TABLE` - 创建表权限   * `DELETE` - 删除数据权限   * `CREATE_SCHEMA` - 创建schema权限
type PermissionActions struct {
	value string
}

type PermissionActionsEnum struct {
	ALL           PermissionActions
	SELECT        PermissionActions
	UPDATE        PermissionActions
	CREATE        PermissionActions
	DROP          PermissionActions
	ALTER         PermissionActions
	INDEX         PermissionActions
	READ          PermissionActions
	WRITE         PermissionActions
	INSERT        PermissionActions
	CREATE_TABLE  PermissionActions
	DELETE        PermissionActions
	CREATE_SCHEMA PermissionActions
}

func GetPermissionActionsEnum() PermissionActionsEnum {
	return PermissionActionsEnum{
		ALL: PermissionActions{
			value: "ALL",
		},
		SELECT: PermissionActions{
			value: "SELECT",
		},
		UPDATE: PermissionActions{
			value: "UPDATE",
		},
		CREATE: PermissionActions{
			value: "CREATE",
		},
		DROP: PermissionActions{
			value: "DROP",
		},
		ALTER: PermissionActions{
			value: "ALTER",
		},
		INDEX: PermissionActions{
			value: "INDEX",
		},
		READ: PermissionActions{
			value: "READ",
		},
		WRITE: PermissionActions{
			value: "WRITE",
		},
		INSERT: PermissionActions{
			value: "INSERT",
		},
		CREATE_TABLE: PermissionActions{
			value: "CREATE_TABLE",
		},
		DELETE: PermissionActions{
			value: "DELETE",
		},
		CREATE_SCHEMA: PermissionActions{
			value: "CREATE_SCHEMA",
		},
	}
}

func (c PermissionActions) Value() string {
	return c.value
}

func (c PermissionActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionActions) UnmarshalJSON(b []byte) error {
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
