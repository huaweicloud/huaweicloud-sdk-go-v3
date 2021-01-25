package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type ConnectionInfo struct {
	Name           *string                       `json:"name,omitempty"`
	ConnectionType *ConnectionInfoConnectionType `json:"connectionType,omitempty"`
	Config         *interface{}                  `json:"config,omitempty"`
}

func (o ConnectionInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConnectionInfo struct{}"
	}

	return strings.Join([]string{"ConnectionInfo", string(data)}, " ")
}

type ConnectionInfoConnectionType struct {
	value string
}

type ConnectionInfoConnectionTypeEnum struct {
	DWS         ConnectionInfoConnectionType
	DLI         ConnectionInfoConnectionType
	SPARK_SQL   ConnectionInfoConnectionType
	HIVE        ConnectionInfoConnectionType
	RDS         ConnectionInfoConnectionType
	CLOUD_TABLE ConnectionInfoConnectionType
}

func GetConnectionInfoConnectionTypeEnum() ConnectionInfoConnectionTypeEnum {
	return ConnectionInfoConnectionTypeEnum{
		DWS: ConnectionInfoConnectionType{
			value: "DWS",
		},
		DLI: ConnectionInfoConnectionType{
			value: "DLI",
		},
		SPARK_SQL: ConnectionInfoConnectionType{
			value: "SparkSQL",
		},
		HIVE: ConnectionInfoConnectionType{
			value: "Hive",
		},
		RDS: ConnectionInfoConnectionType{
			value: "RDS",
		},
		CLOUD_TABLE: ConnectionInfoConnectionType{
			value: "CloudTable",
		},
	}
}

func (c ConnectionInfoConnectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ConnectionInfoConnectionType) UnmarshalJSON(b []byte) error {
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
