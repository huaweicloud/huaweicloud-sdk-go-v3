package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateDatasourceInfoResponse struct {
	// 数据源ID

	DatasourceId *string `json:"datasource_id,omitempty"`
	// 数据源名称

	DatasourceName *string `json:"datasource_name,omitempty"`
	// 数据源类型 - DWS - MYSQL - KAFKA - API - OBS - SAP - MRSHBASE - MRSHDFS - MRSHIVE - WEBSOCKET - SQLSERVER - ORACLE - POSTGRESQL - REDIS - MONGODB - DIS - HL7 - RABBITMQ - SNMP - IBMMQ - CUSTOMIZED (自定义类型) - ACTIVEMQ - ARTEMISMQ - FTP - HIVE - HANA - FIKAFKA - MRSKAFKA - FIHDFS - FIHIVE - GAUSS200 - GAUSS100 - LDAP - DB2 - TAURUS

	DatasourceType *CreateDatasourceInfoResponseDatasourceType `json:"datasource_type,omitempty"`
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

	AppPermission  *[]CreateDatasourceInfoResponseAppPermission `json:"app_permission,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o CreateDatasourceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatasourceInfoResponse struct{}"
	}

	return strings.Join([]string{"CreateDatasourceInfoResponse", string(data)}, " ")
}

type CreateDatasourceInfoResponseDatasourceType struct {
	value string
}

type CreateDatasourceInfoResponseDatasourceTypeEnum struct {
	DWS        CreateDatasourceInfoResponseDatasourceType
	MYSQL      CreateDatasourceInfoResponseDatasourceType
	KAFKA      CreateDatasourceInfoResponseDatasourceType
	API        CreateDatasourceInfoResponseDatasourceType
	OBS        CreateDatasourceInfoResponseDatasourceType
	SAP        CreateDatasourceInfoResponseDatasourceType
	MRSHBASE   CreateDatasourceInfoResponseDatasourceType
	MRSHDFS    CreateDatasourceInfoResponseDatasourceType
	MRSHIVE    CreateDatasourceInfoResponseDatasourceType
	WEBSOCKET  CreateDatasourceInfoResponseDatasourceType
	SQLSERVER  CreateDatasourceInfoResponseDatasourceType
	ORACLE     CreateDatasourceInfoResponseDatasourceType
	POSTGRESQL CreateDatasourceInfoResponseDatasourceType
	REDIS      CreateDatasourceInfoResponseDatasourceType
	MONGODB    CreateDatasourceInfoResponseDatasourceType
	DIS        CreateDatasourceInfoResponseDatasourceType
	HL7        CreateDatasourceInfoResponseDatasourceType
	RABBITMQ   CreateDatasourceInfoResponseDatasourceType
	SNMP       CreateDatasourceInfoResponseDatasourceType
	IBMMQ      CreateDatasourceInfoResponseDatasourceType
	CUSTOMIZED CreateDatasourceInfoResponseDatasourceType
	ACTIVEMQ   CreateDatasourceInfoResponseDatasourceType
	ARTEMISMQ  CreateDatasourceInfoResponseDatasourceType
	FTP        CreateDatasourceInfoResponseDatasourceType
	HIVE       CreateDatasourceInfoResponseDatasourceType
	HANA       CreateDatasourceInfoResponseDatasourceType
	FIKAFKA    CreateDatasourceInfoResponseDatasourceType
	MRSKAFKA   CreateDatasourceInfoResponseDatasourceType
	FIHDFS     CreateDatasourceInfoResponseDatasourceType
	FIHIVE     CreateDatasourceInfoResponseDatasourceType
	GAUSS200   CreateDatasourceInfoResponseDatasourceType
	GAUSS100   CreateDatasourceInfoResponseDatasourceType
	LDAP       CreateDatasourceInfoResponseDatasourceType
	DB2        CreateDatasourceInfoResponseDatasourceType
	TAURUS     CreateDatasourceInfoResponseDatasourceType
}

func GetCreateDatasourceInfoResponseDatasourceTypeEnum() CreateDatasourceInfoResponseDatasourceTypeEnum {
	return CreateDatasourceInfoResponseDatasourceTypeEnum{
		DWS: CreateDatasourceInfoResponseDatasourceType{
			value: "DWS",
		},
		MYSQL: CreateDatasourceInfoResponseDatasourceType{
			value: "MYSQL",
		},
		KAFKA: CreateDatasourceInfoResponseDatasourceType{
			value: "KAFKA",
		},
		API: CreateDatasourceInfoResponseDatasourceType{
			value: "API",
		},
		OBS: CreateDatasourceInfoResponseDatasourceType{
			value: "OBS",
		},
		SAP: CreateDatasourceInfoResponseDatasourceType{
			value: "SAP",
		},
		MRSHBASE: CreateDatasourceInfoResponseDatasourceType{
			value: "MRSHBASE",
		},
		MRSHDFS: CreateDatasourceInfoResponseDatasourceType{
			value: "MRSHDFS",
		},
		MRSHIVE: CreateDatasourceInfoResponseDatasourceType{
			value: "MRSHIVE",
		},
		WEBSOCKET: CreateDatasourceInfoResponseDatasourceType{
			value: "WEBSOCKET",
		},
		SQLSERVER: CreateDatasourceInfoResponseDatasourceType{
			value: "SQLSERVER",
		},
		ORACLE: CreateDatasourceInfoResponseDatasourceType{
			value: "ORACLE",
		},
		POSTGRESQL: CreateDatasourceInfoResponseDatasourceType{
			value: "POSTGRESQL",
		},
		REDIS: CreateDatasourceInfoResponseDatasourceType{
			value: "REDIS",
		},
		MONGODB: CreateDatasourceInfoResponseDatasourceType{
			value: "MONGODB",
		},
		DIS: CreateDatasourceInfoResponseDatasourceType{
			value: "DIS",
		},
		HL7: CreateDatasourceInfoResponseDatasourceType{
			value: "HL7",
		},
		RABBITMQ: CreateDatasourceInfoResponseDatasourceType{
			value: "RABBITMQ",
		},
		SNMP: CreateDatasourceInfoResponseDatasourceType{
			value: "SNMP",
		},
		IBMMQ: CreateDatasourceInfoResponseDatasourceType{
			value: "IBMMQ",
		},
		CUSTOMIZED: CreateDatasourceInfoResponseDatasourceType{
			value: "CUSTOMIZED",
		},
		ACTIVEMQ: CreateDatasourceInfoResponseDatasourceType{
			value: "ACTIVEMQ",
		},
		ARTEMISMQ: CreateDatasourceInfoResponseDatasourceType{
			value: "ARTEMISMQ",
		},
		FTP: CreateDatasourceInfoResponseDatasourceType{
			value: "FTP",
		},
		HIVE: CreateDatasourceInfoResponseDatasourceType{
			value: "HIVE",
		},
		HANA: CreateDatasourceInfoResponseDatasourceType{
			value: "HANA",
		},
		FIKAFKA: CreateDatasourceInfoResponseDatasourceType{
			value: "FIKAFKA",
		},
		MRSKAFKA: CreateDatasourceInfoResponseDatasourceType{
			value: "MRSKAFKA",
		},
		FIHDFS: CreateDatasourceInfoResponseDatasourceType{
			value: "FIHDFS",
		},
		FIHIVE: CreateDatasourceInfoResponseDatasourceType{
			value: "FIHIVE",
		},
		GAUSS200: CreateDatasourceInfoResponseDatasourceType{
			value: "GAUSS200",
		},
		GAUSS100: CreateDatasourceInfoResponseDatasourceType{
			value: "GAUSS100",
		},
		LDAP: CreateDatasourceInfoResponseDatasourceType{
			value: "LDAP",
		},
		DB2: CreateDatasourceInfoResponseDatasourceType{
			value: "DB2",
		},
		TAURUS: CreateDatasourceInfoResponseDatasourceType{
			value: "TAURUS",
		},
	}
}

func (c CreateDatasourceInfoResponseDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDatasourceInfoResponseDatasourceType) UnmarshalJSON(b []byte) error {
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

type CreateDatasourceInfoResponseAppPermission struct {
	value string
}

type CreateDatasourceInfoResponseAppPermissionEnum struct {
	READ   CreateDatasourceInfoResponseAppPermission
	ACCESS CreateDatasourceInfoResponseAppPermission
	DELETE CreateDatasourceInfoResponseAppPermission
	MODIFY CreateDatasourceInfoResponseAppPermission
}

func GetCreateDatasourceInfoResponseAppPermissionEnum() CreateDatasourceInfoResponseAppPermissionEnum {
	return CreateDatasourceInfoResponseAppPermissionEnum{
		READ: CreateDatasourceInfoResponseAppPermission{
			value: "read",
		},
		ACCESS: CreateDatasourceInfoResponseAppPermission{
			value: "access",
		},
		DELETE: CreateDatasourceInfoResponseAppPermission{
			value: "delete",
		},
		MODIFY: CreateDatasourceInfoResponseAppPermission{
			value: "modify",
		},
	}
}

func (c CreateDatasourceInfoResponseAppPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDatasourceInfoResponseAppPermission) UnmarshalJSON(b []byte) error {
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
