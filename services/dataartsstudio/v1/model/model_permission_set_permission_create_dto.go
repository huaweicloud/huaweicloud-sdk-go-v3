package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionSetPermissionCreateDto struct {

	// 数据连接id
	DwId *string `json:"dw_id,omitempty"`

	// 权限类型, DENY, ALLOW
	PermissionType *PermissionSetPermissionCreateDtoPermissionType `json:"permission_type,omitempty"`

	// 权限操作列表
	PermissionActions *[]PermissionSetPermissionCreateDtoPermissionActions `json:"permission_actions,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据源类型, HIVE
	DatasourceType *PermissionSetPermissionCreateDtoDatasourceType `json:"datasource_type,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 模式名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 命名空间。无效参数，待下线。
	Namespace *string `json:"namespace,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 行级策略。无效参数，待下线。
	RowLevelSecurity *string `json:"row_level_security,omitempty"`

	// url路径名称, MRS存算分离或者HIVE指定location场景下使用。
	Url *string `json:"url,omitempty"`
}

func (o PermissionSetPermissionCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSetPermissionCreateDto struct{}"
	}

	return strings.Join([]string{"PermissionSetPermissionCreateDto", string(data)}, " ")
}

type PermissionSetPermissionCreateDtoPermissionType struct {
	value string
}

type PermissionSetPermissionCreateDtoPermissionTypeEnum struct {
	DENY  PermissionSetPermissionCreateDtoPermissionType
	ALLOW PermissionSetPermissionCreateDtoPermissionType
}

func GetPermissionSetPermissionCreateDtoPermissionTypeEnum() PermissionSetPermissionCreateDtoPermissionTypeEnum {
	return PermissionSetPermissionCreateDtoPermissionTypeEnum{
		DENY: PermissionSetPermissionCreateDtoPermissionType{
			value: "DENY",
		},
		ALLOW: PermissionSetPermissionCreateDtoPermissionType{
			value: "ALLOW",
		},
	}
}

func (c PermissionSetPermissionCreateDtoPermissionType) Value() string {
	return c.value
}

func (c PermissionSetPermissionCreateDtoPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetPermissionCreateDtoPermissionType) UnmarshalJSON(b []byte) error {
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

type PermissionSetPermissionCreateDtoPermissionActions struct {
	value string
}

type PermissionSetPermissionCreateDtoPermissionActionsEnum struct {
	ALL    PermissionSetPermissionCreateDtoPermissionActions
	SELECT PermissionSetPermissionCreateDtoPermissionActions
	UPDATE PermissionSetPermissionCreateDtoPermissionActions
	CREATE PermissionSetPermissionCreateDtoPermissionActions
	DROP   PermissionSetPermissionCreateDtoPermissionActions
	ALTER  PermissionSetPermissionCreateDtoPermissionActions
	INDEX  PermissionSetPermissionCreateDtoPermissionActions
	LOCK   PermissionSetPermissionCreateDtoPermissionActions
	READ   PermissionSetPermissionCreateDtoPermissionActions
	WRITE  PermissionSetPermissionCreateDtoPermissionActions
}

func GetPermissionSetPermissionCreateDtoPermissionActionsEnum() PermissionSetPermissionCreateDtoPermissionActionsEnum {
	return PermissionSetPermissionCreateDtoPermissionActionsEnum{
		ALL: PermissionSetPermissionCreateDtoPermissionActions{
			value: "ALL",
		},
		SELECT: PermissionSetPermissionCreateDtoPermissionActions{
			value: "SELECT",
		},
		UPDATE: PermissionSetPermissionCreateDtoPermissionActions{
			value: "UPDATE",
		},
		CREATE: PermissionSetPermissionCreateDtoPermissionActions{
			value: "CREATE",
		},
		DROP: PermissionSetPermissionCreateDtoPermissionActions{
			value: "DROP",
		},
		ALTER: PermissionSetPermissionCreateDtoPermissionActions{
			value: "ALTER",
		},
		INDEX: PermissionSetPermissionCreateDtoPermissionActions{
			value: "INDEX",
		},
		LOCK: PermissionSetPermissionCreateDtoPermissionActions{
			value: "LOCK",
		},
		READ: PermissionSetPermissionCreateDtoPermissionActions{
			value: "READ",
		},
		WRITE: PermissionSetPermissionCreateDtoPermissionActions{
			value: "WRITE",
		},
	}
}

func (c PermissionSetPermissionCreateDtoPermissionActions) Value() string {
	return c.value
}

func (c PermissionSetPermissionCreateDtoPermissionActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetPermissionCreateDtoPermissionActions) UnmarshalJSON(b []byte) error {
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

type PermissionSetPermissionCreateDtoDatasourceType struct {
	value string
}

type PermissionSetPermissionCreateDtoDatasourceTypeEnum struct {
	HIVE PermissionSetPermissionCreateDtoDatasourceType
}

func GetPermissionSetPermissionCreateDtoDatasourceTypeEnum() PermissionSetPermissionCreateDtoDatasourceTypeEnum {
	return PermissionSetPermissionCreateDtoDatasourceTypeEnum{
		HIVE: PermissionSetPermissionCreateDtoDatasourceType{
			value: "HIVE",
		},
	}
}

func (c PermissionSetPermissionCreateDtoDatasourceType) Value() string {
	return c.value
}

func (c PermissionSetPermissionCreateDtoDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetPermissionCreateDtoDatasourceType) UnmarshalJSON(b []byte) error {
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
