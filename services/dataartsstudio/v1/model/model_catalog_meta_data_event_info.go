package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CatalogMetaDataEventInfo 元数据实时同步的事件详细信息
type CatalogMetaDataEventInfo struct {

	// 事件发生时的时间戳
	EventTs *int64 `json:"event_ts,omitempty"`

	// 事件类型
	EventType *CatalogMetaDataEventInfoEventType `json:"event_type,omitempty"`

	// 事件消息，Map<String,Object>结构
	EventMessage *interface{} `json:"event_message,omitempty"`
}

func (o CatalogMetaDataEventInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogMetaDataEventInfo struct{}"
	}

	return strings.Join([]string{"CatalogMetaDataEventInfo", string(data)}, " ")
}

type CatalogMetaDataEventInfoEventType struct {
	value string
}

type CatalogMetaDataEventInfoEventTypeEnum struct {
	ADD_FOREIGN_KEY_EVENT         CatalogMetaDataEventInfoEventType
	ADD_NOT_NULL_CONSTRAINT_EVENT CatalogMetaDataEventInfoEventType
	ADD_PRIMARY_KEY_EVENT         CatalogMetaDataEventInfoEventType
	ADD_UNIQUE_CONSTRAINT_EVENT   CatalogMetaDataEventInfoEventType
	ALTER_DATABASE_EVENT          CatalogMetaDataEventInfoEventType
	ADD_PARTITION_EVENT           CatalogMetaDataEventInfoEventType
	ALTER_PARTITION_EVENT         CatalogMetaDataEventInfoEventType
	ALTER_TABLE_EVENT             CatalogMetaDataEventInfoEventType
	ALTER_CATALOG_EVENT           CatalogMetaDataEventInfoEventType
	CREATE_CATALOG_EVENT          CatalogMetaDataEventInfoEventType
	CREATE_DATABASE_EVENT         CatalogMetaDataEventInfoEventType
	CREATE_FUNCTION_EVENT         CatalogMetaDataEventInfoEventType
	CREATE_TABLE_EVENT            CatalogMetaDataEventInfoEventType
	DROP_CONSTRAINT_EVENT         CatalogMetaDataEventInfoEventType
	DROP_DATABASE_EVENT           CatalogMetaDataEventInfoEventType
	DROP_FUNCTION_EVENT           CatalogMetaDataEventInfoEventType
	DROP_PARTITION_EVENT          CatalogMetaDataEventInfoEventType
	DROP_TABLE_EVENT              CatalogMetaDataEventInfoEventType
	DROP_CATALOG_EVENT            CatalogMetaDataEventInfoEventType
	ADD_INDEX_EVENT               CatalogMetaDataEventInfoEventType
	ALTER_INDEX_EVENT             CatalogMetaDataEventInfoEventType
	DROP_INDEX_EVENT              CatalogMetaDataEventInfoEventType
	ALTER_SCHEMA_EVENT            CatalogMetaDataEventInfoEventType
	CREATE_SCHEMA_EVENT           CatalogMetaDataEventInfoEventType
	DROP_SCHEMA_EVENT             CatalogMetaDataEventInfoEventType
	ALTER_COLUMN_EVENT            CatalogMetaDataEventInfoEventType
	ADD_COLUMN_EVENT              CatalogMetaDataEventInfoEventType
	DROP_COLUMN_EVENT             CatalogMetaDataEventInfoEventType
	ALTER_TRIGGER_EVENT           CatalogMetaDataEventInfoEventType
	ADD_TRIGGER_EVENT             CatalogMetaDataEventInfoEventType
	DROP_TRIGGER_EVENT            CatalogMetaDataEventInfoEventType
}

func GetCatalogMetaDataEventInfoEventTypeEnum() CatalogMetaDataEventInfoEventTypeEnum {
	return CatalogMetaDataEventInfoEventTypeEnum{
		ADD_FOREIGN_KEY_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AddForeignKeyEvent",
		},
		ADD_NOT_NULL_CONSTRAINT_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AddNotNullConstraintEvent",
		},
		ADD_PRIMARY_KEY_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AddPrimaryKeyEvent",
		},
		ADD_UNIQUE_CONSTRAINT_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AddUniqueConstraintEvent",
		},
		ALTER_DATABASE_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AlterDatabaseEvent",
		},
		ADD_PARTITION_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AddPartitionEvent",
		},
		ALTER_PARTITION_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AlterPartitionEvent",
		},
		ALTER_TABLE_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AlterTableEvent",
		},
		ALTER_CATALOG_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AlterCatalogEvent",
		},
		CREATE_CATALOG_EVENT: CatalogMetaDataEventInfoEventType{
			value: "CreateCatalogEvent",
		},
		CREATE_DATABASE_EVENT: CatalogMetaDataEventInfoEventType{
			value: "CreateDatabaseEvent",
		},
		CREATE_FUNCTION_EVENT: CatalogMetaDataEventInfoEventType{
			value: "CreateFunctionEvent",
		},
		CREATE_TABLE_EVENT: CatalogMetaDataEventInfoEventType{
			value: "CreateTableEvent",
		},
		DROP_CONSTRAINT_EVENT: CatalogMetaDataEventInfoEventType{
			value: "DropConstraintEvent",
		},
		DROP_DATABASE_EVENT: CatalogMetaDataEventInfoEventType{
			value: "DropDatabaseEvent",
		},
		DROP_FUNCTION_EVENT: CatalogMetaDataEventInfoEventType{
			value: "DropFunctionEvent",
		},
		DROP_PARTITION_EVENT: CatalogMetaDataEventInfoEventType{
			value: "DropPartitionEvent",
		},
		DROP_TABLE_EVENT: CatalogMetaDataEventInfoEventType{
			value: "DropTableEvent",
		},
		DROP_CATALOG_EVENT: CatalogMetaDataEventInfoEventType{
			value: "DropCatalogEvent",
		},
		ADD_INDEX_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AddIndexEvent",
		},
		ALTER_INDEX_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AlterIndexEvent",
		},
		DROP_INDEX_EVENT: CatalogMetaDataEventInfoEventType{
			value: "DropIndexEvent",
		},
		ALTER_SCHEMA_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AlterSchemaEvent",
		},
		CREATE_SCHEMA_EVENT: CatalogMetaDataEventInfoEventType{
			value: "CreateSchemaEvent",
		},
		DROP_SCHEMA_EVENT: CatalogMetaDataEventInfoEventType{
			value: "DropSchemaEvent",
		},
		ALTER_COLUMN_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AlterColumnEvent",
		},
		ADD_COLUMN_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AddColumnEvent",
		},
		DROP_COLUMN_EVENT: CatalogMetaDataEventInfoEventType{
			value: "DropColumnEvent",
		},
		ALTER_TRIGGER_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AlterTriggerEvent",
		},
		ADD_TRIGGER_EVENT: CatalogMetaDataEventInfoEventType{
			value: "AddTriggerEvent",
		},
		DROP_TRIGGER_EVENT: CatalogMetaDataEventInfoEventType{
			value: "DropTriggerEvent",
		},
	}
}

func (c CatalogMetaDataEventInfoEventType) Value() string {
	return c.value
}

func (c CatalogMetaDataEventInfoEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CatalogMetaDataEventInfoEventType) UnmarshalJSON(b []byte) error {
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
