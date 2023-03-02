package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ClusterDataConnectorMap struct {

	// 数据连接关联ID值
	MapId *int32 `json:"map_id,omitempty"`

	// 数据连接ID值
	ConnectorId *string `json:"connector_id,omitempty"`

	// 组件名
	ComponentName *string `json:"component_name,omitempty"`

	// 组件角色类型。 - hive_metastore：Hive Metastore角色 - hive_data：Hive角色 - hbase_data：Hbase角色 - ranger_data：Ranger角色
	RoleType *ClusterDataConnectorMapRoleType `json:"role_type,omitempty"`

	// 数据连接类型。 - LOCAL_DB：本地元数据 - RDS_POSTGRES：RDS服务PostgreSQL数据库 - RDS_MYSQL：RDS服务MySQL数据库 - gaussdb-mysql：云数据库GaussDB(for MySQL)
	SourceType *ClusterDataConnectorMapSourceType `json:"source_type,omitempty"`

	// 关联集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 数据连接状态。 - 0：代表正常状态 - 1：代表使用中
	Status *int32 `json:"status,omitempty"`
}

func (o ClusterDataConnectorMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDataConnectorMap struct{}"
	}

	return strings.Join([]string{"ClusterDataConnectorMap", string(data)}, " ")
}

type ClusterDataConnectorMapRoleType struct {
	value string
}

type ClusterDataConnectorMapRoleTypeEnum struct {
	LOCAL_DB      ClusterDataConnectorMapRoleType
	RDS_POSTGRES  ClusterDataConnectorMapRoleType
	RDS_MYSQL     ClusterDataConnectorMapRoleType
	GAUSSDB_MYSQL ClusterDataConnectorMapRoleType
}

func GetClusterDataConnectorMapRoleTypeEnum() ClusterDataConnectorMapRoleTypeEnum {
	return ClusterDataConnectorMapRoleTypeEnum{
		LOCAL_DB: ClusterDataConnectorMapRoleType{
			value: "LOCAL_DB",
		},
		RDS_POSTGRES: ClusterDataConnectorMapRoleType{
			value: "RDS_POSTGRES",
		},
		RDS_MYSQL: ClusterDataConnectorMapRoleType{
			value: "RDS_MYSQL",
		},
		GAUSSDB_MYSQL: ClusterDataConnectorMapRoleType{
			value: "gaussdb-mysql",
		},
	}
}

func (c ClusterDataConnectorMapRoleType) Value() string {
	return c.value
}

func (c ClusterDataConnectorMapRoleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterDataConnectorMapRoleType) UnmarshalJSON(b []byte) error {
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

type ClusterDataConnectorMapSourceType struct {
	value string
}

type ClusterDataConnectorMapSourceTypeEnum struct {
	LOCAL_DB      ClusterDataConnectorMapSourceType
	RDS_POSTGRES  ClusterDataConnectorMapSourceType
	RDS_MYSQL     ClusterDataConnectorMapSourceType
	GAUSSDB_MYSQL ClusterDataConnectorMapSourceType
}

func GetClusterDataConnectorMapSourceTypeEnum() ClusterDataConnectorMapSourceTypeEnum {
	return ClusterDataConnectorMapSourceTypeEnum{
		LOCAL_DB: ClusterDataConnectorMapSourceType{
			value: "LOCAL_DB",
		},
		RDS_POSTGRES: ClusterDataConnectorMapSourceType{
			value: "RDS_POSTGRES",
		},
		RDS_MYSQL: ClusterDataConnectorMapSourceType{
			value: "RDS_MYSQL",
		},
		GAUSSDB_MYSQL: ClusterDataConnectorMapSourceType{
			value: "gaussdb-mysql",
		},
	}
}

func (c ClusterDataConnectorMapSourceType) Value() string {
	return c.value
}

func (c ClusterDataConnectorMapSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterDataConnectorMapSourceType) UnmarshalJSON(b []byte) error {
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
