package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type ShowDataourceDetailResponse struct {
	// 数据源ID

	DatasourceId *string `json:"datasource_id,omitempty"`
	// 数据源名称

	DatasourceName *string `json:"datasource_name,omitempty"`
	// 数据源类型 - DWS - MYSQL - KAFKA - API - OBS - SAP - MRSHBASE - MRSHDFS - MRSHIVE - WEBSOCKET - SQLSERVER - ORACLE - POSTGRESQL - REDIS - MONGODB - DIS - HL7 - RABBITMQ - SNMP - IBMMQ - CUSTOMIZED (自定义类型) - ACTIVEMQ - ARTEMISMQ - FTP - HIVE - HANA - FIKAFKA - MRSKAFKA - FIHDFS - FIHIVE - GAUSS200 - GAUSS100 - LDAP - DB2 - TAURUS

	DatasourceType *ShowDataourceDetailResponseDatasourceType `json:"datasource_type,omitempty"`
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

	AppPermission  *[]ShowDataourceDetailResponseAppPermission `json:"app_permission,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o ShowDataourceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataourceDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDataourceDetailResponse", string(data)}, " ")
}

type ShowDataourceDetailResponseDatasourceType struct {
	value string
}

type ShowDataourceDetailResponseDatasourceTypeEnum struct {
	DWS        ShowDataourceDetailResponseDatasourceType
	MYSQL      ShowDataourceDetailResponseDatasourceType
	KAFKA      ShowDataourceDetailResponseDatasourceType
	API        ShowDataourceDetailResponseDatasourceType
	OBS        ShowDataourceDetailResponseDatasourceType
	SAP        ShowDataourceDetailResponseDatasourceType
	MRSHBASE   ShowDataourceDetailResponseDatasourceType
	MRSHDFS    ShowDataourceDetailResponseDatasourceType
	MRSHIVE    ShowDataourceDetailResponseDatasourceType
	WEBSOCKET  ShowDataourceDetailResponseDatasourceType
	SQLSERVER  ShowDataourceDetailResponseDatasourceType
	ORACLE     ShowDataourceDetailResponseDatasourceType
	POSTGRESQL ShowDataourceDetailResponseDatasourceType
	REDIS      ShowDataourceDetailResponseDatasourceType
	MONGODB    ShowDataourceDetailResponseDatasourceType
	DIS        ShowDataourceDetailResponseDatasourceType
	HL7        ShowDataourceDetailResponseDatasourceType
	RABBITMQ   ShowDataourceDetailResponseDatasourceType
	SNMP       ShowDataourceDetailResponseDatasourceType
	IBMMQ      ShowDataourceDetailResponseDatasourceType
	CUSTOMIZED ShowDataourceDetailResponseDatasourceType
	ACTIVEMQ   ShowDataourceDetailResponseDatasourceType
	ARTEMISMQ  ShowDataourceDetailResponseDatasourceType
	FTP        ShowDataourceDetailResponseDatasourceType
	HIVE       ShowDataourceDetailResponseDatasourceType
	HANA       ShowDataourceDetailResponseDatasourceType
	FIKAFKA    ShowDataourceDetailResponseDatasourceType
	MRSKAFKA   ShowDataourceDetailResponseDatasourceType
	FIHDFS     ShowDataourceDetailResponseDatasourceType
	FIHIVE     ShowDataourceDetailResponseDatasourceType
	GAUSS200   ShowDataourceDetailResponseDatasourceType
	GAUSS100   ShowDataourceDetailResponseDatasourceType
	LDAP       ShowDataourceDetailResponseDatasourceType
	DB2        ShowDataourceDetailResponseDatasourceType
	TAURUS     ShowDataourceDetailResponseDatasourceType
}

func GetShowDataourceDetailResponseDatasourceTypeEnum() ShowDataourceDetailResponseDatasourceTypeEnum {
	return ShowDataourceDetailResponseDatasourceTypeEnum{
		DWS: ShowDataourceDetailResponseDatasourceType{
			value: "DWS",
		},
		MYSQL: ShowDataourceDetailResponseDatasourceType{
			value: "MYSQL",
		},
		KAFKA: ShowDataourceDetailResponseDatasourceType{
			value: "KAFKA",
		},
		API: ShowDataourceDetailResponseDatasourceType{
			value: "API",
		},
		OBS: ShowDataourceDetailResponseDatasourceType{
			value: "OBS",
		},
		SAP: ShowDataourceDetailResponseDatasourceType{
			value: "SAP",
		},
		MRSHBASE: ShowDataourceDetailResponseDatasourceType{
			value: "MRSHBASE",
		},
		MRSHDFS: ShowDataourceDetailResponseDatasourceType{
			value: "MRSHDFS",
		},
		MRSHIVE: ShowDataourceDetailResponseDatasourceType{
			value: "MRSHIVE",
		},
		WEBSOCKET: ShowDataourceDetailResponseDatasourceType{
			value: "WEBSOCKET",
		},
		SQLSERVER: ShowDataourceDetailResponseDatasourceType{
			value: "SQLSERVER",
		},
		ORACLE: ShowDataourceDetailResponseDatasourceType{
			value: "ORACLE",
		},
		POSTGRESQL: ShowDataourceDetailResponseDatasourceType{
			value: "POSTGRESQL",
		},
		REDIS: ShowDataourceDetailResponseDatasourceType{
			value: "REDIS",
		},
		MONGODB: ShowDataourceDetailResponseDatasourceType{
			value: "MONGODB",
		},
		DIS: ShowDataourceDetailResponseDatasourceType{
			value: "DIS",
		},
		HL7: ShowDataourceDetailResponseDatasourceType{
			value: "HL7",
		},
		RABBITMQ: ShowDataourceDetailResponseDatasourceType{
			value: "RABBITMQ",
		},
		SNMP: ShowDataourceDetailResponseDatasourceType{
			value: "SNMP",
		},
		IBMMQ: ShowDataourceDetailResponseDatasourceType{
			value: "IBMMQ",
		},
		CUSTOMIZED: ShowDataourceDetailResponseDatasourceType{
			value: "CUSTOMIZED",
		},
		ACTIVEMQ: ShowDataourceDetailResponseDatasourceType{
			value: "ACTIVEMQ",
		},
		ARTEMISMQ: ShowDataourceDetailResponseDatasourceType{
			value: "ARTEMISMQ",
		},
		FTP: ShowDataourceDetailResponseDatasourceType{
			value: "FTP",
		},
		HIVE: ShowDataourceDetailResponseDatasourceType{
			value: "HIVE",
		},
		HANA: ShowDataourceDetailResponseDatasourceType{
			value: "HANA",
		},
		FIKAFKA: ShowDataourceDetailResponseDatasourceType{
			value: "FIKAFKA",
		},
		MRSKAFKA: ShowDataourceDetailResponseDatasourceType{
			value: "MRSKAFKA",
		},
		FIHDFS: ShowDataourceDetailResponseDatasourceType{
			value: "FIHDFS",
		},
		FIHIVE: ShowDataourceDetailResponseDatasourceType{
			value: "FIHIVE",
		},
		GAUSS200: ShowDataourceDetailResponseDatasourceType{
			value: "GAUSS200",
		},
		GAUSS100: ShowDataourceDetailResponseDatasourceType{
			value: "GAUSS100",
		},
		LDAP: ShowDataourceDetailResponseDatasourceType{
			value: "LDAP",
		},
		DB2: ShowDataourceDetailResponseDatasourceType{
			value: "DB2",
		},
		TAURUS: ShowDataourceDetailResponseDatasourceType{
			value: "TAURUS",
		},
	}
}

func (c ShowDataourceDetailResponseDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDataourceDetailResponseDatasourceType) UnmarshalJSON(b []byte) error {
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

type ShowDataourceDetailResponseAppPermission struct {
	value string
}

type ShowDataourceDetailResponseAppPermissionEnum struct {
	READ   ShowDataourceDetailResponseAppPermission
	ACCESS ShowDataourceDetailResponseAppPermission
	DELETE ShowDataourceDetailResponseAppPermission
	MODIFY ShowDataourceDetailResponseAppPermission
}

func GetShowDataourceDetailResponseAppPermissionEnum() ShowDataourceDetailResponseAppPermissionEnum {
	return ShowDataourceDetailResponseAppPermissionEnum{
		READ: ShowDataourceDetailResponseAppPermission{
			value: "read",
		},
		ACCESS: ShowDataourceDetailResponseAppPermission{
			value: "access",
		},
		DELETE: ShowDataourceDetailResponseAppPermission{
			value: "delete",
		},
		MODIFY: ShowDataourceDetailResponseAppPermission{
			value: "modify",
		},
	}
}

func (c ShowDataourceDetailResponseAppPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDataourceDetailResponseAppPermission) UnmarshalJSON(b []byte) error {
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
