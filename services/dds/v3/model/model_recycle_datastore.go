package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RecycleDatastore 数据库信息。
type RecycleDatastore struct {

	// 数据库版本类型。取值“DDS-Community”。
	Type RecycleDatastoreType `json:"type"`

	// 数据库版本。支持3.4及以上版本。取值为“3.4”、“4.0”、“4.2”、“4.4”或“5.0”。
	Version string `json:"version"`
}

func (o RecycleDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleDatastore struct{}"
	}

	return strings.Join([]string{"RecycleDatastore", string(data)}, " ")
}

type RecycleDatastoreType struct {
	value string
}

type RecycleDatastoreTypeEnum struct {
	DDS_COMMUNITY RecycleDatastoreType
}

func GetRecycleDatastoreTypeEnum() RecycleDatastoreTypeEnum {
	return RecycleDatastoreTypeEnum{
		DDS_COMMUNITY: RecycleDatastoreType{
			value: "DDS-Community",
		},
	}
}

func (c RecycleDatastoreType) Value() string {
	return c.value
}

func (c RecycleDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleDatastoreType) UnmarshalJSON(b []byte) error {
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
