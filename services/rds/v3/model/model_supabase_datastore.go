package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SupabaseDatastore 数据库信息。
type SupabaseDatastore struct {

	// 数据库引擎。
	Type SupabaseDatastoreType `json:"type"`

	// 数据库版本。
	Version string `json:"version"`
}

func (o SupabaseDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseDatastore struct{}"
	}

	return strings.Join([]string{"SupabaseDatastore", string(data)}, " ")
}

type SupabaseDatastoreType struct {
	value string
}

type SupabaseDatastoreTypeEnum struct {
	POSTGRE_SQL SupabaseDatastoreType
}

func GetSupabaseDatastoreTypeEnum() SupabaseDatastoreTypeEnum {
	return SupabaseDatastoreTypeEnum{
		POSTGRE_SQL: SupabaseDatastoreType{
			value: "PostgreSQL",
		},
	}
}

func (c SupabaseDatastoreType) Value() string {
	return c.value
}

func (c SupabaseDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SupabaseDatastoreType) UnmarshalJSON(b []byte) error {
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
