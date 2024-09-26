package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Datastore 数据库信息。
type Datastore struct {

	// 数据库版本类型。取值“DDS-Community”。
	Type DatastoreType `json:"type"`

	// 数据库版本。支持3.4及以上版本。取值为“3.4”、“4.0”、“4.2”、“4.4”、“5.0”。
	Version string `json:"version"`

	// 存储引擎。支持WiredTiger存储引擎。取值为“wiredTiger”。
	StorageEngine DatastoreStorageEngine `json:"storage_engine"`
}

func (o Datastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datastore struct{}"
	}

	return strings.Join([]string{"Datastore", string(data)}, " ")
}

type DatastoreType struct {
	value string
}

type DatastoreTypeEnum struct {
	DDS_COMMUNITY DatastoreType
}

func GetDatastoreTypeEnum() DatastoreTypeEnum {
	return DatastoreTypeEnum{
		DDS_COMMUNITY: DatastoreType{
			value: "DDS-Community",
		},
	}
}

func (c DatastoreType) Value() string {
	return c.value
}

func (c DatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatastoreType) UnmarshalJSON(b []byte) error {
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

type DatastoreStorageEngine struct {
	value string
}

type DatastoreStorageEngineEnum struct {
	WIRED_TIGER DatastoreStorageEngine
	ROCKS_DB    DatastoreStorageEngine
}

func GetDatastoreStorageEngineEnum() DatastoreStorageEngineEnum {
	return DatastoreStorageEngineEnum{
		WIRED_TIGER: DatastoreStorageEngine{
			value: "wiredTiger",
		},
		ROCKS_DB: DatastoreStorageEngine{
			value: "rocksDB",
		},
	}
}

func (c DatastoreStorageEngine) Value() string {
	return c.value
}

func (c DatastoreStorageEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatastoreStorageEngine) UnmarshalJSON(b []byte) error {
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
