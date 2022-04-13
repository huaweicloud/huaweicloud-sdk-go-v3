package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 数据源
type DatasourceInfo struct {
	// 数据源名称，数据源名称不能包含&、<、>、\"、'、(、) ，长度为1~255字符

	DatasourceName *string `json:"datasource_name,omitempty"`
	// 数据源类型 - DWS - MYSQL - KAFKA - API - OBS - SAP - MRSHBASE - MRSHDFS - MRSHIVE - WEBSOCKET - SQLSERVER - ORACLE - POSTGRESQL - REDIS - MONGODB - DIS - HL7 - RABBITMQ - SNMP - IBMMQ - CUSTOMIZED (自定义类型) - ACTIVEMQ - ARTEMISMQ - FTP - HIVE - HANA - FIKAFKA - MRSKAFKA - FIHDFS - FIHIVE - GAUSS200 - GAUSS100 - LDAP - DB2 - TAURUS

	DatasourceType *DatasourceInfoDatasourceType `json:"datasource_type,omitempty"`
	// 数据源所属应用ID

	AppId *string `json:"app_id,omitempty"`
	// 数据源所属连接器Id，自定义数据源必填

	CustomPluginId *string `json:"custom_plugin_id,omitempty"`

	Content *Content `json:"content,omitempty"`
	// 数据源描述

	Description *string `json:"description,omitempty"`
}

func (o DatasourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatasourceInfo struct{}"
	}

	return strings.Join([]string{"DatasourceInfo", string(data)}, " ")
}

type DatasourceInfoDatasourceType struct {
	value string
}

type DatasourceInfoDatasourceTypeEnum struct {
	DWS        DatasourceInfoDatasourceType
	MYSQL      DatasourceInfoDatasourceType
	KAFKA      DatasourceInfoDatasourceType
	API        DatasourceInfoDatasourceType
	OBS        DatasourceInfoDatasourceType
	SAP        DatasourceInfoDatasourceType
	MRSHBASE   DatasourceInfoDatasourceType
	MRSHDFS    DatasourceInfoDatasourceType
	MRSHIVE    DatasourceInfoDatasourceType
	WEBSOCKET  DatasourceInfoDatasourceType
	SQLSERVER  DatasourceInfoDatasourceType
	ORACLE     DatasourceInfoDatasourceType
	POSTGRESQL DatasourceInfoDatasourceType
	REDIS      DatasourceInfoDatasourceType
	MONGODB    DatasourceInfoDatasourceType
	DIS        DatasourceInfoDatasourceType
	HL7        DatasourceInfoDatasourceType
	RABBITMQ   DatasourceInfoDatasourceType
	SNMP       DatasourceInfoDatasourceType
	IBMMQ      DatasourceInfoDatasourceType
	CUSTOMIZED DatasourceInfoDatasourceType
	ACTIVEMQ   DatasourceInfoDatasourceType
	ARTEMISMQ  DatasourceInfoDatasourceType
	FTP        DatasourceInfoDatasourceType
	HIVE       DatasourceInfoDatasourceType
	HANA       DatasourceInfoDatasourceType
	FIKAFKA    DatasourceInfoDatasourceType
	MRSKAFKA   DatasourceInfoDatasourceType
	FIHDFS     DatasourceInfoDatasourceType
	FIHIVE     DatasourceInfoDatasourceType
	GAUSS200   DatasourceInfoDatasourceType
	GAUSS100   DatasourceInfoDatasourceType
	LDAP       DatasourceInfoDatasourceType
	DB2        DatasourceInfoDatasourceType
	TAURUS     DatasourceInfoDatasourceType
}

func GetDatasourceInfoDatasourceTypeEnum() DatasourceInfoDatasourceTypeEnum {
	return DatasourceInfoDatasourceTypeEnum{
		DWS: DatasourceInfoDatasourceType{
			value: "DWS",
		},
		MYSQL: DatasourceInfoDatasourceType{
			value: "MYSQL",
		},
		KAFKA: DatasourceInfoDatasourceType{
			value: "KAFKA",
		},
		API: DatasourceInfoDatasourceType{
			value: "API",
		},
		OBS: DatasourceInfoDatasourceType{
			value: "OBS",
		},
		SAP: DatasourceInfoDatasourceType{
			value: "SAP",
		},
		MRSHBASE: DatasourceInfoDatasourceType{
			value: "MRSHBASE",
		},
		MRSHDFS: DatasourceInfoDatasourceType{
			value: "MRSHDFS",
		},
		MRSHIVE: DatasourceInfoDatasourceType{
			value: "MRSHIVE",
		},
		WEBSOCKET: DatasourceInfoDatasourceType{
			value: "WEBSOCKET",
		},
		SQLSERVER: DatasourceInfoDatasourceType{
			value: "SQLSERVER",
		},
		ORACLE: DatasourceInfoDatasourceType{
			value: "ORACLE",
		},
		POSTGRESQL: DatasourceInfoDatasourceType{
			value: "POSTGRESQL",
		},
		REDIS: DatasourceInfoDatasourceType{
			value: "REDIS",
		},
		MONGODB: DatasourceInfoDatasourceType{
			value: "MONGODB",
		},
		DIS: DatasourceInfoDatasourceType{
			value: "DIS",
		},
		HL7: DatasourceInfoDatasourceType{
			value: "HL7",
		},
		RABBITMQ: DatasourceInfoDatasourceType{
			value: "RABBITMQ",
		},
		SNMP: DatasourceInfoDatasourceType{
			value: "SNMP",
		},
		IBMMQ: DatasourceInfoDatasourceType{
			value: "IBMMQ",
		},
		CUSTOMIZED: DatasourceInfoDatasourceType{
			value: "CUSTOMIZED",
		},
		ACTIVEMQ: DatasourceInfoDatasourceType{
			value: "ACTIVEMQ",
		},
		ARTEMISMQ: DatasourceInfoDatasourceType{
			value: "ARTEMISMQ",
		},
		FTP: DatasourceInfoDatasourceType{
			value: "FTP",
		},
		HIVE: DatasourceInfoDatasourceType{
			value: "HIVE",
		},
		HANA: DatasourceInfoDatasourceType{
			value: "HANA",
		},
		FIKAFKA: DatasourceInfoDatasourceType{
			value: "FIKAFKA",
		},
		MRSKAFKA: DatasourceInfoDatasourceType{
			value: "MRSKAFKA",
		},
		FIHDFS: DatasourceInfoDatasourceType{
			value: "FIHDFS",
		},
		FIHIVE: DatasourceInfoDatasourceType{
			value: "FIHIVE",
		},
		GAUSS200: DatasourceInfoDatasourceType{
			value: "GAUSS200",
		},
		GAUSS100: DatasourceInfoDatasourceType{
			value: "GAUSS100",
		},
		LDAP: DatasourceInfoDatasourceType{
			value: "LDAP",
		},
		DB2: DatasourceInfoDatasourceType{
			value: "DB2",
		},
		TAURUS: DatasourceInfoDatasourceType{
			value: "TAURUS",
		},
	}
}

func (c DatasourceInfoDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatasourceInfoDatasourceType) UnmarshalJSON(b []byte) error {
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
