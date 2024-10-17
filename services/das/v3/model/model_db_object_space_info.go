package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DbObjectSpaceInfo 数据库对象空间信息，数据来源于information_schema.tables系统表。已用空间包含数据空间、索引空间和碎片空间。
type DbObjectSpaceInfo struct {

	// 对象类型，如果是table，同时需要传database_id
	ObjectType DbObjectSpaceInfoObjectType `json:"object_type"`

	// 对象名称
	ObjectName string `json:"object_name"`

	// 对象ID
	ObjectId *string `json:"object_id,omitempty"`

	// 已使用空间，以字节为单位
	UsedSize *int64 `json:"used_size,omitempty"`

	// 数据空间，以字节为单位
	DataSize *int64 `json:"data_size,omitempty"`

	// 索引空间，以字节为单位
	IndexSize *int64 `json:"index_size,omitempty"`

	// 碎片空间，以字节为单位
	FreeSize *int64 `json:"free_size,omitempty"`

	// 碎片率
	FreeRate *float64 `json:"free_rate,omitempty"`

	// 估算值行数，以字节为单位
	EstimatedRows *int64 `json:"estimated_rows,omitempty"`
}

func (o DbObjectSpaceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbObjectSpaceInfo struct{}"
	}

	return strings.Join([]string{"DbObjectSpaceInfo", string(data)}, " ")
}

type DbObjectSpaceInfoObjectType struct {
	value string
}

type DbObjectSpaceInfoObjectTypeEnum struct {
	DATABASE DbObjectSpaceInfoObjectType
	TABLE    DbObjectSpaceInfoObjectType
}

func GetDbObjectSpaceInfoObjectTypeEnum() DbObjectSpaceInfoObjectTypeEnum {
	return DbObjectSpaceInfoObjectTypeEnum{
		DATABASE: DbObjectSpaceInfoObjectType{
			value: "database",
		},
		TABLE: DbObjectSpaceInfoObjectType{
			value: "table",
		},
	}
}

func (c DbObjectSpaceInfoObjectType) Value() string {
	return c.value
}

func (c DbObjectSpaceInfoObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DbObjectSpaceInfoObjectType) UnmarshalJSON(b []byte) error {
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
