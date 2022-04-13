package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListDatasourcesRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 每页显示条目数量，最大数量999，超过999后只返回999

	Limit *int32 `json:"limit,omitempty"`
	// 分页查询，分页的偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int32 `json:"offset,omitempty"`
	// 数据源类型 - DWS - MYSQL - KAFKA - API - OBS - SAP - MRSHBASE - MRSHDFS - MRSHIVE - WEBSOCKET - SQLSERVER - ORACLE - POSTGRESQL - REDIS - MONGODB - DIS - HL7 - RABBITMQ - SNMP - IBMMQ - CUSTOMIZED (自定义类型) - ACTIVEMQ - ARTEMISMQ - FTP - HIVE - HANA - FIKAFKA - MRSKAFKA - FIHDFS - FIHIVE - GAUSS200 - GAUSS100 - LDAP - DB2 - TAURUS

	DatasourceType *ListDatasourcesRequestDatasourceType `json:"datasource_type,omitempty"`
	// 排序字段（CREATED_DATE）

	SortField *string `json:"sort_field,omitempty"`
	// 查询数据源排序的类型，增序还是降序，可为空

	SortType *ListDatasourcesRequestSortType `json:"sort_type,omitempty"`
	// 数据源名称,支持模糊匹配

	Name *string `json:"name,omitempty"`
	// 集成应用ID

	AppId *string `json:"app_id,omitempty"`
	// 连接器ID

	CustomPluginId *string `json:"custom_plugin_id,omitempty"`
}

func (o ListDatasourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasourcesRequest struct{}"
	}

	return strings.Join([]string{"ListDatasourcesRequest", string(data)}, " ")
}

type ListDatasourcesRequestDatasourceType struct {
	value string
}

type ListDatasourcesRequestDatasourceTypeEnum struct {
	DWS        ListDatasourcesRequestDatasourceType
	MYSQL      ListDatasourcesRequestDatasourceType
	KAFKA      ListDatasourcesRequestDatasourceType
	API        ListDatasourcesRequestDatasourceType
	OBS        ListDatasourcesRequestDatasourceType
	SAP        ListDatasourcesRequestDatasourceType
	MRSHBASE   ListDatasourcesRequestDatasourceType
	MRSHDFS    ListDatasourcesRequestDatasourceType
	MRSHIVE    ListDatasourcesRequestDatasourceType
	WEBSOCKET  ListDatasourcesRequestDatasourceType
	SQLSERVER  ListDatasourcesRequestDatasourceType
	ORACLE     ListDatasourcesRequestDatasourceType
	POSTGRESQL ListDatasourcesRequestDatasourceType
	REDIS      ListDatasourcesRequestDatasourceType
	MONGODB    ListDatasourcesRequestDatasourceType
	DIS        ListDatasourcesRequestDatasourceType
	HL7        ListDatasourcesRequestDatasourceType
	RABBITMQ   ListDatasourcesRequestDatasourceType
	SNMP       ListDatasourcesRequestDatasourceType
	IBMMQ      ListDatasourcesRequestDatasourceType
	CUSTOMIZED ListDatasourcesRequestDatasourceType
	ACTIVEMQ   ListDatasourcesRequestDatasourceType
	ARTEMISMQ  ListDatasourcesRequestDatasourceType
	FTP        ListDatasourcesRequestDatasourceType
	HIVE       ListDatasourcesRequestDatasourceType
	HANA       ListDatasourcesRequestDatasourceType
	FIKAFKA    ListDatasourcesRequestDatasourceType
	MRSKAFKA   ListDatasourcesRequestDatasourceType
	FIHDFS     ListDatasourcesRequestDatasourceType
	FIHIVE     ListDatasourcesRequestDatasourceType
	GAUSS200   ListDatasourcesRequestDatasourceType
	GAUSS100   ListDatasourcesRequestDatasourceType
	LDAP       ListDatasourcesRequestDatasourceType
	DB2        ListDatasourcesRequestDatasourceType
	TAURUS     ListDatasourcesRequestDatasourceType
}

func GetListDatasourcesRequestDatasourceTypeEnum() ListDatasourcesRequestDatasourceTypeEnum {
	return ListDatasourcesRequestDatasourceTypeEnum{
		DWS: ListDatasourcesRequestDatasourceType{
			value: "DWS",
		},
		MYSQL: ListDatasourcesRequestDatasourceType{
			value: "MYSQL",
		},
		KAFKA: ListDatasourcesRequestDatasourceType{
			value: "KAFKA",
		},
		API: ListDatasourcesRequestDatasourceType{
			value: "API",
		},
		OBS: ListDatasourcesRequestDatasourceType{
			value: "OBS",
		},
		SAP: ListDatasourcesRequestDatasourceType{
			value: "SAP",
		},
		MRSHBASE: ListDatasourcesRequestDatasourceType{
			value: "MRSHBASE",
		},
		MRSHDFS: ListDatasourcesRequestDatasourceType{
			value: "MRSHDFS",
		},
		MRSHIVE: ListDatasourcesRequestDatasourceType{
			value: "MRSHIVE",
		},
		WEBSOCKET: ListDatasourcesRequestDatasourceType{
			value: "WEBSOCKET",
		},
		SQLSERVER: ListDatasourcesRequestDatasourceType{
			value: "SQLSERVER",
		},
		ORACLE: ListDatasourcesRequestDatasourceType{
			value: "ORACLE",
		},
		POSTGRESQL: ListDatasourcesRequestDatasourceType{
			value: "POSTGRESQL",
		},
		REDIS: ListDatasourcesRequestDatasourceType{
			value: "REDIS",
		},
		MONGODB: ListDatasourcesRequestDatasourceType{
			value: "MONGODB",
		},
		DIS: ListDatasourcesRequestDatasourceType{
			value: "DIS",
		},
		HL7: ListDatasourcesRequestDatasourceType{
			value: "HL7",
		},
		RABBITMQ: ListDatasourcesRequestDatasourceType{
			value: "RABBITMQ",
		},
		SNMP: ListDatasourcesRequestDatasourceType{
			value: "SNMP",
		},
		IBMMQ: ListDatasourcesRequestDatasourceType{
			value: "IBMMQ",
		},
		CUSTOMIZED: ListDatasourcesRequestDatasourceType{
			value: "CUSTOMIZED",
		},
		ACTIVEMQ: ListDatasourcesRequestDatasourceType{
			value: "ACTIVEMQ",
		},
		ARTEMISMQ: ListDatasourcesRequestDatasourceType{
			value: "ARTEMISMQ",
		},
		FTP: ListDatasourcesRequestDatasourceType{
			value: "FTP",
		},
		HIVE: ListDatasourcesRequestDatasourceType{
			value: "HIVE",
		},
		HANA: ListDatasourcesRequestDatasourceType{
			value: "HANA",
		},
		FIKAFKA: ListDatasourcesRequestDatasourceType{
			value: "FIKAFKA",
		},
		MRSKAFKA: ListDatasourcesRequestDatasourceType{
			value: "MRSKAFKA",
		},
		FIHDFS: ListDatasourcesRequestDatasourceType{
			value: "FIHDFS",
		},
		FIHIVE: ListDatasourcesRequestDatasourceType{
			value: "FIHIVE",
		},
		GAUSS200: ListDatasourcesRequestDatasourceType{
			value: "GAUSS200",
		},
		GAUSS100: ListDatasourcesRequestDatasourceType{
			value: "GAUSS100",
		},
		LDAP: ListDatasourcesRequestDatasourceType{
			value: "LDAP",
		},
		DB2: ListDatasourcesRequestDatasourceType{
			value: "DB2",
		},
		TAURUS: ListDatasourcesRequestDatasourceType{
			value: "TAURUS",
		},
	}
}

func (c ListDatasourcesRequestDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatasourcesRequestDatasourceType) UnmarshalJSON(b []byte) error {
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

type ListDatasourcesRequestSortType struct {
	value string
}

type ListDatasourcesRequestSortTypeEnum struct {
	ASC  ListDatasourcesRequestSortType
	DESC ListDatasourcesRequestSortType
}

func GetListDatasourcesRequestSortTypeEnum() ListDatasourcesRequestSortTypeEnum {
	return ListDatasourcesRequestSortTypeEnum{
		ASC: ListDatasourcesRequestSortType{
			value: "ASC",
		},
		DESC: ListDatasourcesRequestSortType{
			value: "DESC",
		},
	}
}

func (c ListDatasourcesRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatasourcesRequestSortType) UnmarshalJSON(b []byte) error {
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
