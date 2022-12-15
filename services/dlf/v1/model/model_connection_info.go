package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConnectionInfo struct {
	Name string `json:"name"`

	Type ConnectionInfoType `json:"type"`

	Config *interface{} `json:"config,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o ConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionInfo struct{}"
	}

	return strings.Join([]string{"ConnectionInfo", string(data)}, " ")
}

type ConnectionInfoType struct {
	value string
}

type ConnectionInfoTypeEnum struct {
	DWS         ConnectionInfoType
	DLI         ConnectionInfoType
	SPARK_SQL   ConnectionInfoType
	HIVE        ConnectionInfoType
	RDS         ConnectionInfoType
	CLOUD_TABLE ConnectionInfoType
}

func GetConnectionInfoTypeEnum() ConnectionInfoTypeEnum {
	return ConnectionInfoTypeEnum{
		DWS: ConnectionInfoType{
			value: "DWS",
		},
		DLI: ConnectionInfoType{
			value: "DLI",
		},
		SPARK_SQL: ConnectionInfoType{
			value: "SparkSQL",
		},
		HIVE: ConnectionInfoType{
			value: "Hive",
		},
		RDS: ConnectionInfoType{
			value: "RDS",
		},
		CLOUD_TABLE: ConnectionInfoType{
			value: "CloudTable",
		},
	}
}

func (c ConnectionInfoType) Value() string {
	return c.value
}

func (c ConnectionInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionInfoType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
