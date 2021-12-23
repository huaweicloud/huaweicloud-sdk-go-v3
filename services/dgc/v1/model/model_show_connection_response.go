package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowConnectionResponse struct {
	Name *string `json:"name,omitempty"`

	ConnectionType *ShowConnectionResponseConnectionType `json:"connectionType,omitempty"`

	Config         *interface{} `json:"config,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectionResponse", string(data)}, " ")
}

type ShowConnectionResponseConnectionType struct {
	value string
}

type ShowConnectionResponseConnectionTypeEnum struct {
	DWS         ShowConnectionResponseConnectionType
	DLI         ShowConnectionResponseConnectionType
	SPARK_SQL   ShowConnectionResponseConnectionType
	HIVE        ShowConnectionResponseConnectionType
	RDS         ShowConnectionResponseConnectionType
	CLOUD_TABLE ShowConnectionResponseConnectionType
}

func GetShowConnectionResponseConnectionTypeEnum() ShowConnectionResponseConnectionTypeEnum {
	return ShowConnectionResponseConnectionTypeEnum{
		DWS: ShowConnectionResponseConnectionType{
			value: "DWS",
		},
		DLI: ShowConnectionResponseConnectionType{
			value: "DLI",
		},
		SPARK_SQL: ShowConnectionResponseConnectionType{
			value: "SparkSQL",
		},
		HIVE: ShowConnectionResponseConnectionType{
			value: "Hive",
		},
		RDS: ShowConnectionResponseConnectionType{
			value: "RDS",
		},
		CLOUD_TABLE: ShowConnectionResponseConnectionType{
			value: "CloudTable",
		},
	}
}

func (c ShowConnectionResponseConnectionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConnectionResponseConnectionType) UnmarshalJSON(b []byte) error {
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
