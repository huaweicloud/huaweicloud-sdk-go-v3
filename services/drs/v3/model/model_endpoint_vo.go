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
	Id *string `json:"id,omitempty"`

	// 对象id。
	ObjId *string `json:"obj_id,omitempty"`

	// RDS实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 数据库类型
	DbType *EndpointVoDbType `json:"db_type,omitempty"`

	// 数据库用户。
	DbUser *string `json:"db_user,omitempty"`

	// 数据库密码。
	DbPassword *string `json:"db_password,omitempty"`

	// 管理IP。
	ManageIp *string `json:"manage_ip,omitempty"`

	// 流量IP。
	TrafficIp *string `json:"traffic_ip,omitempty"`

	// 数据库端口。
	DbPort *int32 `json:"db_port,omitempty"`

	// RDS实例所在region。
	Region *string `json:"region,omitempty"`

	// 创建日期，格式yyyy-MM-dd'T'HH:mm:ss'Z'
	CreatedAt *string `json:"created_at,omitempty"`

	// 修改日期，格式yyyy-MM-dd'T'HH:mm:ss'Z'
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 迁移实例所在的私有IP。
	Ip *string `json:"ip,omitempty"`

	// 迁移实例所在的公网IP。
	PublicIp *string `json:"public_ip,omitempty"`

	// 可用区azCode。
	AzCode *string `json:"az_code,omitempty"`

	// 源库所在的安全组id。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 源库所在的子网id。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 源库所在的虚拟私有云id。
	VpcId *string `json:"vpc_id,omitempty"`

	// 迁移实例的磁盘大小。
	VolumeSize *int64 `json:"volume_size,omitempty"`

	// 全量迁移用户密码，密文。
	FullTransUserPwd *string `json:"full_trans_user_pwd,omitempty"`

	// 增量迁移用户密码，密文。
	IncrementTransUserPwd *string `json:"increment_trans_user_pwd,omitempty"`

	// 是否SSL安全连接。
	SslLink *bool `json:"ssl_link,omitempty"`

	// SSL证书内容。
	SslCertKey *string `json:"ssl_cert_key,omitempty"`

	// SSL证书名字。
	SslCertName *string `json:"ssl_cert_name,omitempty"`

	// SSL证书内容checksum值。
	SslCertCheckSum *string `json:"ssl_cert_check_sum,omitempty"`

	// SSL证书密码，密文。
	SslCertPassword *string `json:"ssl_cert_password,omitempty"`

	// 数据库版本。
	DbVersion *string `json:"db_version,omitempty"`

	// mongoHa模式。 - Sharding 集群 - ReplicaSet 副本集 - ReplicaSingle 单节点
	MongoHaMode *EndpointVoMongoHaMode `json:"mongo_ha_mode,omitempty"`

	// RDS实例projectId。
	ProjectId *string `json:"project_id,omitempty"`

	// 集群模式。 - Single：单节点RDS - Ha：主备RDS - GR：金融版RDS - Sharding：mongodb 集群或DDM的模式，均默认为分片 - Sharding4.0+：mongodb 集群版本4.0+，默认为不分片 - ReplicaSet：mongodb 副本集,Replica RDS只读副本 - ReplicaSingle：mongodb 单节点 - Cluster：集群 - Independent：gaussdbv5 independent模式 - Combined：gaussdbv5 Combined模式 - Distributed ：分布式taurus - NoSharding：非集群模式
	ClusterMode *EndpointVoClusterMode `json:"cluster_mode,omitempty"`

	// RDS实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// Oracle服务名serviceName。
	DbName *string `json:"db_name,omitempty"`

	// mrskafka topic名称。
	Topic *string `json:"topic,omitempty"`

	// MRSkafka是否开启kerberos认证 - 0非安全认证 - 1安全认证
	SafeMode *int32 `json:"safe_mode,omitempty"`

	KerberosVo *KerberosVo `json:"kerberos_vo,omitempty"`

	// 多写数据库Id。
	MultiWriteDbId *string `json:"multi_write_db_id,omitempty"`
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
			value: "Sharding",
		},
		REPLICA_SET: EndpointVoMongoHaMode{
			value: "ReplicaSet",
		},
		REPLICA_SINGLE: EndpointVoMongoHaMode{
			value: "ReplicaSingle",
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
	SINGLE         EndpointVoClusterMode
	HA             EndpointVoClusterMode
	GR             EndpointVoClusterMode
	SHARDING       EndpointVoClusterMode
	SHARDING4_0    EndpointVoClusterMode
	REPLICA_SET    EndpointVoClusterMode
	REPLICA        EndpointVoClusterMode
	REPLICA_SINGLE EndpointVoClusterMode
	CLUSTER        EndpointVoClusterMode
	INDEPENDENT    EndpointVoClusterMode
	COMBINED       EndpointVoClusterMode
	DISTRIBUTED    EndpointVoClusterMode
	NO_SHARDING    EndpointVoClusterMode
}

func GetEndpointVoClusterModeEnum() EndpointVoClusterModeEnum {
	return EndpointVoClusterModeEnum{
		SINGLE: EndpointVoClusterMode{
			value: "Single",
		},
		HA: EndpointVoClusterMode{
			value: "Ha",
		},
		GR: EndpointVoClusterMode{
			value: "GR",
		},
		SHARDING: EndpointVoClusterMode{
			value: "Sharding",
		},
		SHARDING4_0: EndpointVoClusterMode{
			value: "Sharding4.0+",
		},
		REPLICA_SET: EndpointVoClusterMode{
			value: "ReplicaSet",
		},
		REPLICA: EndpointVoClusterMode{
			value: "Replica",
		},
		REPLICA_SINGLE: EndpointVoClusterMode{
			value: "ReplicaSingle",
		},
		CLUSTER: EndpointVoClusterMode{
			value: "Cluster",
		},
		INDEPENDENT: EndpointVoClusterMode{
			value: "Independent",
		},
		COMBINED: EndpointVoClusterMode{
			value: "Combined",
		},
		DISTRIBUTED: EndpointVoClusterMode{
			value: "Distributed",
		},
		NO_SHARDING: EndpointVoClusterMode{
			value: "NoSharding",
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
