package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 数据库信息
type EndpointVo struct {

	// 数据库id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 对象id。
	ObjId *string `json:"obj_id,omitempty" xml:"obj_id"`

	// RDS实例名称。
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name"`

	// 数据库类型
	DbType *EndpointVoDbType `json:"db_type,omitempty" xml:"db_type"`

	// 数据库用户。
	DbUser *string `json:"db_user,omitempty" xml:"db_user"`

	// 数据库密码。
	DbPassword *string `json:"db_password,omitempty" xml:"db_password"`

	// 管理IP。
	ManageIp *string `json:"manage_ip,omitempty" xml:"manage_ip"`

	// 流量IP。
	TrafficIp *string `json:"traffic_ip,omitempty" xml:"traffic_ip"`

	// 数据库端口。
	DbPort *int32 `json:"db_port,omitempty" xml:"db_port"`

	// RDS实例所在region。
	Region *string `json:"region,omitempty" xml:"region"`

	// 创建日期，格式yyyy-MM-dd'T'HH:mm:ss'Z'
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 修改日期，格式yyyy-MM-dd'T'HH:mm:ss'Z'
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 迁移实例所在的私有IP。
	Ip *string `json:"ip,omitempty" xml:"ip"`

	// 迁移实例所在的公网IP。
	PublicIp *string `json:"public_ip,omitempty" xml:"public_ip"`

	// 可用区azCode。
	AzCode *string `json:"az_code,omitempty" xml:"az_code"`

	// 源库所在的安全组id。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// 源库所在的子网id。
	SubnetId *string `json:"subnet_id,omitempty" xml:"subnet_id"`

	// 源库所在的虚拟私有云id。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 迁移实例的磁盘大小。
	VolumeSize *int64 `json:"volume_size,omitempty" xml:"volume_size"`

	// 全量迁移用户密码，密文。
	FullTransUserPwd *string `json:"full_trans_user_pwd,omitempty" xml:"full_trans_user_pwd"`

	// 增量迁移用户密码，密文。
	IncrementTransUserPwd *string `json:"increment_trans_user_pwd,omitempty" xml:"increment_trans_user_pwd"`

	// 是否SSL安全连接。
	SslLink *bool `json:"ssl_link,omitempty" xml:"ssl_link"`

	// SSL证书内容。
	SslCertKey *string `json:"ssl_cert_key,omitempty" xml:"ssl_cert_key"`

	// SSL证书名字。
	SslCertName *string `json:"ssl_cert_name,omitempty" xml:"ssl_cert_name"`

	// SSL证书内容checksum值。
	SslCertCheckSum *string `json:"ssl_cert_check_sum,omitempty" xml:"ssl_cert_check_sum"`

	// SSL证书密码，密文。
	SslCertPassword *string `json:"ssl_cert_password,omitempty" xml:"ssl_cert_password"`

	// 数据库版本。
	DbVersion *string `json:"db_version,omitempty" xml:"db_version"`

	// mongoHa模式
	MongoHaMode *EndpointVoMongoHaMode `json:"mongo_ha_mode,omitempty" xml:"mongo_ha_mode"`

	// RDS实例projectId。
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 集群模式
	ClusterMode *EndpointVoClusterMode `json:"cluster_mode,omitempty" xml:"cluster_mode"`

	// RDS实例id。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// Oracle服务名serviceName。
	DbName *string `json:"db_name,omitempty" xml:"db_name"`

	// mrskafka topic名称。
	Topic *string `json:"topic,omitempty" xml:"topic"`

	// MRSkafka是否开启kerberos认证 - 0非安全认证 - 1安全认证
	SafeMode *int32 `json:"safe_mode,omitempty" xml:"safe_mode"`

	KerberosVo *KerberosVo `json:"kerberos_vo,omitempty" xml:"kerberos_vo"`

	// 多写数据库Id。
	MultiWriteDbId *string `json:"multi_write_db_id,omitempty" xml:"multi_write_db_id"`
}

func (o EndpointVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointVo struct{}"
	}

	return strings.Join([]string{"EndpointVo", string(data)}, " ")
}

type EndpointVoDbType struct {
	value string
}

type EndpointVoDbTypeEnum struct {
	MYSQL   EndpointVoDbType
	MONGODB EndpointVoDbType
}

func GetEndpointVoDbTypeEnum() EndpointVoDbTypeEnum {
	return EndpointVoDbTypeEnum{
		MYSQL: EndpointVoDbType{
			value: "mysql",
		},
		MONGODB: EndpointVoDbType{
			value: "mongodb",
		},
	}
}

func (c EndpointVoDbType) Value() string {
	return c.value
}

func (c EndpointVoDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointVoDbType) UnmarshalJSON(b []byte) error {
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

type EndpointVoMongoHaMode struct {
	value string
}

type EndpointVoMongoHaModeEnum struct {
	SHARDING       EndpointVoMongoHaMode
	REPLICA_SET    EndpointVoMongoHaMode
	REPLICA_SINGLE EndpointVoMongoHaMode
}

func GetEndpointVoMongoHaModeEnum() EndpointVoMongoHaModeEnum {
	return EndpointVoMongoHaModeEnum{
		SHARDING: EndpointVoMongoHaMode{
			value: "Sharding 集群",
		},
		REPLICA_SET: EndpointVoMongoHaMode{
			value: "ReplicaSet 副本集",
		},
		REPLICA_SINGLE: EndpointVoMongoHaMode{
			value: "ReplicaSingle 单节点",
		},
	}
}

func (c EndpointVoMongoHaMode) Value() string {
	return c.value
}

func (c EndpointVoMongoHaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointVoMongoHaMode) UnmarshalJSON(b []byte) error {
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

type EndpointVoClusterMode struct {
	value string
}

type EndpointVoClusterModeEnum struct {
	SINGLE_RDS                        EndpointVoClusterMode
	HA_RDS                            EndpointVoClusterMode
	GR_RDS                            EndpointVoClusterMode
	SHARDING_MONGODB_DDM              EndpointVoClusterMode
	REPLICA_SET_MONGODB               EndpointVoClusterMode
	REPLICA_RDS                       EndpointVoClusterMode
	REPLICA_SINGLE_MONGODB            EndpointVoClusterMode
	CLUSTER                           EndpointVoClusterMode
	INDEPENDENT_GAUSSDBV5_INDEPENDENT EndpointVoClusterMode
	COMBINED_GAUSSDBV5_COMBINED       EndpointVoClusterMode
	DISTRIBUTED_TAURUS                EndpointVoClusterMode
}

func GetEndpointVoClusterModeEnum() EndpointVoClusterModeEnum {
	return EndpointVoClusterModeEnum{
		SINGLE_RDS: EndpointVoClusterMode{
			value: "Single 单节点RDS",
		},
		HA_RDS: EndpointVoClusterMode{
			value: "Ha 主备RDS",
		},
		GR_RDS: EndpointVoClusterMode{
			value: "GR 金融版RDS",
		},
		SHARDING_MONGODB_DDM: EndpointVoClusterMode{
			value: "Sharding mongodb 集群或DDM的模式，均默认为分片",
		},
		REPLICA_SET_MONGODB: EndpointVoClusterMode{
			value: "ReplicaSet mongodb 副本集",
		},
		REPLICA_RDS: EndpointVoClusterMode{
			value: "Replica RDS只读副本",
		},
		REPLICA_SINGLE_MONGODB: EndpointVoClusterMode{
			value: "ReplicaSingle mongodb 单节点",
		},
		CLUSTER: EndpointVoClusterMode{
			value: "Cluster 集群",
		},
		INDEPENDENT_GAUSSDBV5_INDEPENDENT: EndpointVoClusterMode{
			value: "Independent gaussdbv5 independent模式",
		},
		COMBINED_GAUSSDBV5_COMBINED: EndpointVoClusterMode{
			value: "Combined gaussdbv5 Combined模式",
		},
		DISTRIBUTED_TAURUS: EndpointVoClusterMode{
			value: "Distributed 分布式taurus",
		},
	}
}

func (c EndpointVoClusterMode) Value() string {
	return c.value
}

func (c EndpointVoClusterMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointVoClusterMode) UnmarshalJSON(b []byte) error {
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
