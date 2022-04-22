package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 数据源
type DataSourceRsp struct {

	// 数据源ID
	DatasourceId *string `json:"datasource_id,omitempty"`

	// 数据源名称
	DatasourceName *string `json:"datasource_name,omitempty"`

	// 数据源类型 - DWS - MYSQL - KAFKA - API - OBS - SAP - MRSHBASE - MRSHDFS - MRSHIVE - WEBSOCKET - SQLSERVER - ORACLE - POSTGRESQL - REDIS - MONGODB - DIS - HL7 - RABBITMQ - SNMP - IBMMQ - CUSTOMIZED (自定义类型) - ACTIVEMQ - ARTEMISMQ - FTP - HIVE - HANA - FIKAFKA - MRSKAFKA - FIHDFS - FIHIVE - GAUSS200 - GAUSS100 - LDAP - DB2 - TAURUS
	DatasourceType *DataSourceRspDatasourceType `json:"datasource_type,omitempty"`

	// 数据源所属虚拟私有云VpcId
	VpcId *string `json:"vpc_id,omitempty"`

	// 数据源所属应用ID
	AppId *string `json:"app_id,omitempty"`

	// 数据源所属应用名称
	AppName *string `json:"app_name,omitempty"`

	// 数据源所属实例Id
	InstanceId *string `json:"instance_id,omitempty"`

	// 数据源创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 数据源修改时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 数据源所属连接器Id
	CustomPluginId *string `json:"custom_plugin_id,omitempty"`

	Content *Content `json:"content,omitempty"`

	// 数据源描述
	Description *string `json:"description,omitempty"`

	// 集成应用权限信息 - read (读权限) - access (调用权限) - delete (删除权限) - modify (修改权限)
	AppPermission *[]DataSourceRspAppPermission `json:"app_permission,omitempty"`
}

func (o DataSourceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataSourceRsp struct{}"
	}

	return strings.Join([]string{"DataSourceRsp", string(data)}, " ")
}

type DataSourceRspDatasourceType struct {
	value string
}

type DataSourceRspDatasourceTypeEnum struct {
	DWS        DataSourceRspDatasourceType
	MYSQL      DataSourceRspDatasourceType
	KAFKA      DataSourceRspDatasourceType
	API        DataSourceRspDatasourceType
	OBS        DataSourceRspDatasourceType
	SAP        DataSourceRspDatasourceType
	MRSHBASE   DataSourceRspDatasourceType
	MRSHDFS    DataSourceRspDatasourceType
	MRSHIVE    DataSourceRspDatasourceType
	WEBSOCKET  DataSourceRspDatasourceType
	SQLSERVER  DataSourceRspDatasourceType
	ORACLE     DataSourceRspDatasourceType
	POSTGRESQL DataSourceRspDatasourceType
	REDIS      DataSourceRspDatasourceType
	MONGODB    DataSourceRspDatasourceType
	DIS        DataSourceRspDatasourceType
	HL7        DataSourceRspDatasourceType
	RABBITMQ   DataSourceRspDatasourceType
	SNMP       DataSourceRspDatasourceType
	IBMMQ      DataSourceRspDatasourceType
	CUSTOMIZED DataSourceRspDatasourceType
	ACTIVEMQ   DataSourceRspDatasourceType
	ARTEMISMQ  DataSourceRspDatasourceType
	FTP        DataSourceRspDatasourceType
	HIVE       DataSourceRspDatasourceType
	HANA       DataSourceRspDatasourceType
	FIKAFKA    DataSourceRspDatasourceType
	MRSKAFKA   DataSourceRspDatasourceType
	FIHDFS     DataSourceRspDatasourceType
	FIHIVE     DataSourceRspDatasourceType
	GAUSS200   DataSourceRspDatasourceType
	GAUSS100   DataSourceRspDatasourceType
	LDAP       DataSourceRspDatasourceType
	DB2        DataSourceRspDatasourceType
	TAURUS     DataSourceRspDatasourceType
}

func GetDataSourceRspDatasourceTypeEnum() DataSourceRspDatasourceTypeEnum {
	return DataSourceRspDatasourceTypeEnum{
		DWS: DataSourceRspDatasourceType{
			value: "DWS",
		},
		MYSQL: DataSourceRspDatasourceType{
			value: "MYSQL",
		},
		KAFKA: DataSourceRspDatasourceType{
			value: "KAFKA",
		},
		API: DataSourceRspDatasourceType{
			value: "API",
		},
		OBS: DataSourceRspDatasourceType{
			value: "OBS",
		},
		SAP: DataSourceRspDatasourceType{
			value: "SAP",
		},
		MRSHBASE: DataSourceRspDatasourceType{
			value: "MRSHBASE",
		},
		MRSHDFS: DataSourceRspDatasourceType{
			value: "MRSHDFS",
		},
		MRSHIVE: DataSourceRspDatasourceType{
			value: "MRSHIVE",
		},
		WEBSOCKET: DataSourceRspDatasourceType{
			value: "WEBSOCKET",
		},
		SQLSERVER: DataSourceRspDatasourceType{
			value: "SQLSERVER",
		},
		ORACLE: DataSourceRspDatasourceType{
			value: "ORACLE",
		},
		POSTGRESQL: DataSourceRspDatasourceType{
			value: "POSTGRESQL",
		},
		REDIS: DataSourceRspDatasourceType{
			value: "REDIS",
		},
		MONGODB: DataSourceRspDatasourceType{
			value: "MONGODB",
		},
		DIS: DataSourceRspDatasourceType{
			value: "DIS",
		},
		HL7: DataSourceRspDatasourceType{
			value: "HL7",
		},
		RABBITMQ: DataSourceRspDatasourceType{
			value: "RABBITMQ",
		},
		SNMP: DataSourceRspDatasourceType{
			value: "SNMP",
		},
		IBMMQ: DataSourceRspDatasourceType{
			value: "IBMMQ",
		},
		CUSTOMIZED: DataSourceRspDatasourceType{
			value: "CUSTOMIZED",
		},
		ACTIVEMQ: DataSourceRspDatasourceType{
			value: "ACTIVEMQ",
		},
		ARTEMISMQ: DataSourceRspDatasourceType{
			value: "ARTEMISMQ",
		},
		FTP: DataSourceRspDatasourceType{
			value: "FTP",
		},
		HIVE: DataSourceRspDatasourceType{
			value: "HIVE",
		},
		HANA: DataSourceRspDatasourceType{
			value: "HANA",
		},
		FIKAFKA: DataSourceRspDatasourceType{
			value: "FIKAFKA",
		},
		MRSKAFKA: DataSourceRspDatasourceType{
			value: "MRSKAFKA",
		},
		FIHDFS: DataSourceRspDatasourceType{
			value: "FIHDFS",
		},
		FIHIVE: DataSourceRspDatasourceType{
			value: "FIHIVE",
		},
		GAUSS200: DataSourceRspDatasourceType{
			value: "GAUSS200",
		},
		GAUSS100: DataSourceRspDatasourceType{
			value: "GAUSS100",
		},
		LDAP: DataSourceRspDatasourceType{
			value: "LDAP",
		},
		DB2: DataSourceRspDatasourceType{
			value: "DB2",
		},
		TAURUS: DataSourceRspDatasourceType{
			value: "TAURUS",
		},
	}
}

func (c DataSourceRspDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataSourceRspDatasourceType) UnmarshalJSON(b []byte) error {
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

type DataSourceRspAppPermission struct {
	value string
}

type DataSourceRspAppPermissionEnum struct {
	READ   DataSourceRspAppPermission
	ACCESS DataSourceRspAppPermission
	DELETE DataSourceRspAppPermission
	MODIFY DataSourceRspAppPermission
}

func GetDataSourceRspAppPermissionEnum() DataSourceRspAppPermissionEnum {
	return DataSourceRspAppPermissionEnum{
		READ: DataSourceRspAppPermission{
			value: "read",
		},
		ACCESS: DataSourceRspAppPermission{
			value: "access",
		},
		DELETE: DataSourceRspAppPermission{
			value: "delete",
		},
		MODIFY: DataSourceRspAppPermission{
			value: "modify",
		},
	}
}

func (c DataSourceRspAppPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataSourceRspAppPermission) UnmarshalJSON(b []byte) error {
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
