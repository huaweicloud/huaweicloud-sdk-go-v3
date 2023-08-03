package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataConnector struct {

	// 数据连接名称。
	ConnectorName string `json:"connector_name"`

	// 数据连接类型。 - RDS_POSTGRES：RDS服务PostgreSQL数据库 - RDS_MYSQL：RDS服务MySQL数据库 - gaussdb-mysql：云数据库GaussDB(for MySQL)
	SourceType DataConnectorSourceType `json:"source_type"`

	// 数据源信息，为json格式，不同数据连接有不同的信息，各数据源的source_info请求内容可参见示例。
	SourceInfo string `json:"source_info"`
}

func (o DataConnector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataConnector struct{}"
	}

	return strings.Join([]string{"DataConnector", string(data)}, " ")
}

type DataConnectorSourceType struct {
	value string
}

type DataConnectorSourceTypeEnum struct {
	RDS_MYSQL     DataConnectorSourceType
	RDS_POSTGRES  DataConnectorSourceType
	GAUSSDB_MYSQL DataConnectorSourceType
}

func GetDataConnectorSourceTypeEnum() DataConnectorSourceTypeEnum {
	return DataConnectorSourceTypeEnum{
		RDS_MYSQL: DataConnectorSourceType{
			value: "RDS_MYSQL",
		},
		RDS_POSTGRES: DataConnectorSourceType{
			value: "RDS_POSTGRES",
		},
		GAUSSDB_MYSQL: DataConnectorSourceType{
			value: "gaussdb-mysql",
		},
	}
}

func (c DataConnectorSourceType) Value() string {
	return c.value
}

func (c DataConnectorSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataConnectorSourceType) UnmarshalJSON(b []byte) error {
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
