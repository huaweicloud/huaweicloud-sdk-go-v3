package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateDatasourceInfoResponse struct {
	// 数据源ID

	DatasourceId *string `json:"datasource_id,omitempty"`
	// 数据源名称

	DatasourceName *string `json:"datasource_name,omitempty"`
	// 数据源类型 - DWS - MYSQL - KAFKA - API - OBS - SAP - MRSHBASE - MRSHDFS - MRSHIVE - WEBSOCKET - SQLSERVER - ORACLE - POSTGRESQL - REDIS - MONGODB - DIS - HL7 - RABBITMQ - SNMP - IBMMQ - CUSTOMIZED (自定义类型) - ACTIVEMQ - ARTEMISMQ - FTP - HIVE - HANA - FIKAFKA - MRSKAFKA - FIHDFS - FIHIVE - GAUSS200 - GAUSS100 - LDAP - DB2 - TAURUS

	DatasourceType *UpdateDatasourceInfoResponseDatasourceType `json:"datasource_type,omitempty"`
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

	AppPermission  *[]UpdateDatasourceInfoResponseAppPermission `json:"app_permission,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o UpdateDatasourceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatasourceInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateDatasourceInfoResponse", string(data)}, " ")
}

type UpdateDatasourceInfoResponseDatasourceType struct {
	value string
}

type UpdateDatasourceInfoResponseDatasourceTypeEnum struct {
	DWS        UpdateDatasourceInfoResponseDatasourceType
	MYSQL      UpdateDatasourceInfoResponseDatasourceType
	KAFKA      UpdateDatasourceInfoResponseDatasourceType
	API        UpdateDatasourceInfoResponseDatasourceType
	OBS        UpdateDatasourceInfoResponseDatasourceType
	SAP        UpdateDatasourceInfoResponseDatasourceType
	MRSHBASE   UpdateDatasourceInfoResponseDatasourceType
	MRSHDFS    UpdateDatasourceInfoResponseDatasourceType
	MRSHIVE    UpdateDatasourceInfoResponseDatasourceType
	WEBSOCKET  UpdateDatasourceInfoResponseDatasourceType
	SQLSERVER  UpdateDatasourceInfoResponseDatasourceType
	ORACLE     UpdateDatasourceInfoResponseDatasourceType
	POSTGRESQL UpdateDatasourceInfoResponseDatasourceType
	REDIS      UpdateDatasourceInfoResponseDatasourceType
	MONGODB    UpdateDatasourceInfoResponseDatasourceType
	DIS        UpdateDatasourceInfoResponseDatasourceType
	HL7        UpdateDatasourceInfoResponseDatasourceType
	RABBITMQ   UpdateDatasourceInfoResponseDatasourceType
	SNMP       UpdateDatasourceInfoResponseDatasourceType
	IBMMQ      UpdateDatasourceInfoResponseDatasourceType
	CUSTOMIZED UpdateDatasourceInfoResponseDatasourceType
	ACTIVEMQ   UpdateDatasourceInfoResponseDatasourceType
	ARTEMISMQ  UpdateDatasourceInfoResponseDatasourceType
	FTP        UpdateDatasourceInfoResponseDatasourceType
	HIVE       UpdateDatasourceInfoResponseDatasourceType
	HANA       UpdateDatasourceInfoResponseDatasourceType
	FIKAFKA    UpdateDatasourceInfoResponseDatasourceType
	MRSKAFKA   UpdateDatasourceInfoResponseDatasourceType
	FIHDFS     UpdateDatasourceInfoResponseDatasourceType
	FIHIVE     UpdateDatasourceInfoResponseDatasourceType
	GAUSS200   UpdateDatasourceInfoResponseDatasourceType
	GAUSS100   UpdateDatasourceInfoResponseDatasourceType
	LDAP       UpdateDatasourceInfoResponseDatasourceType
	DB2        UpdateDatasourceInfoResponseDatasourceType
	TAURUS     UpdateDatasourceInfoResponseDatasourceType
}

func GetUpdateDatasourceInfoResponseDatasourceTypeEnum() UpdateDatasourceInfoResponseDatasourceTypeEnum {
	return UpdateDatasourceInfoResponseDatasourceTypeEnum{
		DWS: UpdateDatasourceInfoResponseDatasourceType{
			value: "DWS",
		},
		MYSQL: UpdateDatasourceInfoResponseDatasourceType{
			value: "MYSQL",
		},
		KAFKA: UpdateDatasourceInfoResponseDatasourceType{
			value: "KAFKA",
		},
		API: UpdateDatasourceInfoResponseDatasourceType{
			value: "API",
		},
		OBS: UpdateDatasourceInfoResponseDatasourceType{
			value: "OBS",
		},
		SAP: UpdateDatasourceInfoResponseDatasourceType{
			value: "SAP",
		},
		MRSHBASE: UpdateDatasourceInfoResponseDatasourceType{
			value: "MRSHBASE",
		},
		MRSHDFS: UpdateDatasourceInfoResponseDatasourceType{
			value: "MRSHDFS",
		},
		MRSHIVE: UpdateDatasourceInfoResponseDatasourceType{
			value: "MRSHIVE",
		},
		WEBSOCKET: UpdateDatasourceInfoResponseDatasourceType{
			value: "WEBSOCKET",
		},
		SQLSERVER: UpdateDatasourceInfoResponseDatasourceType{
			value: "SQLSERVER",
		},
		ORACLE: UpdateDatasourceInfoResponseDatasourceType{
			value: "ORACLE",
		},
		POSTGRESQL: UpdateDatasourceInfoResponseDatasourceType{
			value: "POSTGRESQL",
		},
		REDIS: UpdateDatasourceInfoResponseDatasourceType{
			value: "REDIS",
		},
		MONGODB: UpdateDatasourceInfoResponseDatasourceType{
			value: "MONGODB",
		},
		DIS: UpdateDatasourceInfoResponseDatasourceType{
			value: "DIS",
		},
		HL7: UpdateDatasourceInfoResponseDatasourceType{
			value: "HL7",
		},
		RABBITMQ: UpdateDatasourceInfoResponseDatasourceType{
			value: "RABBITMQ",
		},
		SNMP: UpdateDatasourceInfoResponseDatasourceType{
			value: "SNMP",
		},
		IBMMQ: UpdateDatasourceInfoResponseDatasourceType{
			value: "IBMMQ",
		},
		CUSTOMIZED: UpdateDatasourceInfoResponseDatasourceType{
			value: "CUSTOMIZED",
		},
		ACTIVEMQ: UpdateDatasourceInfoResponseDatasourceType{
			value: "ACTIVEMQ",
		},
		ARTEMISMQ: UpdateDatasourceInfoResponseDatasourceType{
			value: "ARTEMISMQ",
		},
		FTP: UpdateDatasourceInfoResponseDatasourceType{
			value: "FTP",
		},
		HIVE: UpdateDatasourceInfoResponseDatasourceType{
			value: "HIVE",
		},
		HANA: UpdateDatasourceInfoResponseDatasourceType{
			value: "HANA",
		},
		FIKAFKA: UpdateDatasourceInfoResponseDatasourceType{
			value: "FIKAFKA",
		},
		MRSKAFKA: UpdateDatasourceInfoResponseDatasourceType{
			value: "MRSKAFKA",
		},
		FIHDFS: UpdateDatasourceInfoResponseDatasourceType{
			value: "FIHDFS",
		},
		FIHIVE: UpdateDatasourceInfoResponseDatasourceType{
			value: "FIHIVE",
		},
		GAUSS200: UpdateDatasourceInfoResponseDatasourceType{
			value: "GAUSS200",
		},
		GAUSS100: UpdateDatasourceInfoResponseDatasourceType{
			value: "GAUSS100",
		},
		LDAP: UpdateDatasourceInfoResponseDatasourceType{
			value: "LDAP",
		},
		DB2: UpdateDatasourceInfoResponseDatasourceType{
			value: "DB2",
		},
		TAURUS: UpdateDatasourceInfoResponseDatasourceType{
			value: "TAURUS",
		},
	}
}

func (c UpdateDatasourceInfoResponseDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDatasourceInfoResponseDatasourceType) UnmarshalJSON(b []byte) error {
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

type UpdateDatasourceInfoResponseAppPermission struct {
	value string
}

type UpdateDatasourceInfoResponseAppPermissionEnum struct {
	READ   UpdateDatasourceInfoResponseAppPermission
	ACCESS UpdateDatasourceInfoResponseAppPermission
	DELETE UpdateDatasourceInfoResponseAppPermission
	MODIFY UpdateDatasourceInfoResponseAppPermission
}

func GetUpdateDatasourceInfoResponseAppPermissionEnum() UpdateDatasourceInfoResponseAppPermissionEnum {
	return UpdateDatasourceInfoResponseAppPermissionEnum{
		READ: UpdateDatasourceInfoResponseAppPermission{
			value: "read",
		},
		ACCESS: UpdateDatasourceInfoResponseAppPermission{
			value: "access",
		},
		DELETE: UpdateDatasourceInfoResponseAppPermission{
			value: "delete",
		},
		MODIFY: UpdateDatasourceInfoResponseAppPermission{
			value: "modify",
		},
	}
}

func (c UpdateDatasourceInfoResponseAppPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDatasourceInfoResponseAppPermission) UnmarshalJSON(b []byte) error {
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
