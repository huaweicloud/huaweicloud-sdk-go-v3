package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConnectionParam struct {
	Name *string `json:"name,omitempty"`

	ConnectionType *ConnectionParamConnectionType `json:"connectionType,omitempty"`

	Params *interface{} `json:"params,omitempty"`
}

func (o ConnectionParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionParam struct{}"
	}

	return strings.Join([]string{"ConnectionParam", string(data)}, " ")
}

type ConnectionParamConnectionType struct {
	value string
}

type ConnectionParamConnectionTypeEnum struct {
	DWS         ConnectionParamConnectionType
	DLI         ConnectionParamConnectionType
	SPARK_SQL   ConnectionParamConnectionType
	HIVE        ConnectionParamConnectionType
	RDS         ConnectionParamConnectionType
	CLOUD_TABLE ConnectionParamConnectionType
}

func GetConnectionParamConnectionTypeEnum() ConnectionParamConnectionTypeEnum {
	return ConnectionParamConnectionTypeEnum{
		DWS: ConnectionParamConnectionType{
			value: "DWS",
		},
		DLI: ConnectionParamConnectionType{
			value: "DLI",
		},
		SPARK_SQL: ConnectionParamConnectionType{
			value: "SparkSQL",
		},
		HIVE: ConnectionParamConnectionType{
			value: "Hive",
		},
		RDS: ConnectionParamConnectionType{
			value: "RDS",
		},
		CLOUD_TABLE: ConnectionParamConnectionType{
			value: "CloudTable",
		},
	}
}

func (c ConnectionParamConnectionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionParamConnectionType) UnmarshalJSON(b []byte) error {
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
