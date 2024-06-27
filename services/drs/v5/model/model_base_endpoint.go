package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BaseEndpoint 数据库基本信息体。
type BaseEndpoint struct {

	// 数据库信息ID。
	Id *string `json:"id,omitempty"`

	// 数据库场景类型。取值： - oracle：云下自建Oracle数据库。 - ecs_oracle：华为云ECS自建Oracle数据库。 - cloud_gaussdbv5：华为云数据库GaussDB分布式。 - mysql：他云/本地自建MySQL数据库。 - ecs_mysql：华为云ECS自建MySQL数据库。 - cloud_mysql：华为云数据库RDS for MySQL。 - redis：云下自建Redis数据。 - ecs_redis：华为云ECS自建Redis数据。 - rediscluster：云下自建Redis集群数据库。 - ecs_rediscluster：华为云ECS自建Redis集群数据库。 - cloud_gaussdb_redis：华为云数据库GeminiDB Redis。 - postgresql: 云下自建PostgreSQL数据库。 - ecs_postgresql: 华为云ECS自建PostgreSQL数据库。 - cloud_postgresql: 华为云数据库RDS for PostgreSQL。 - mongodb: 云下自建MongoDB数据库。 - ecs_mongodb: 华为云ECS自建MongoDB数据库。 - cloud_mongodb: 华为云文档数据库服务DDS。
	EndpointName BaseEndpointEndpointName `json:"endpoint_name"`

	// 数据库IP。 约束： - 数据库为自建MongoDB时，数据库IP与端口之间用“:”英文冒号拼接，多个值之间请用“,”英文逗号隔开，最多支持填写3个IP地址或域名。 - 数据库为DDS实例时，数据库IP与端口之间用“:”英文冒号拼接，多个IP端口之间请用“,”英文逗号分隔。 - 数据库为Redis集群时，请填写源端Redis集群所有分片的IP地址和对应端口，数据库IP与端口之间用“:”英文冒号拼接，多个IP端口之间请用“,”英文逗号分隔，并且推荐填写集群分片的Slave节点的IP地址。最多支持填写32个IP地址或域名，多个值之间请用英文逗号隔开。 示例： - MySQL：ip - MongoDB：ip:port,ip:port,ip:port - DDS：ip:port,ip:port  - Redis集群：ip:port,ip:port
	Ip *string `json:"ip,omitempty"`

	// 数据库端口。  约束：输入范围为1-65535之间的整数。
	DbPort *string `json:"db_port,omitempty"`

	// 数据库用户名。
	DbUser string `json:"db_user"`

	// 数据库密码。
	DbPassword string `json:"db_password"`

	// 华为云数据库实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 华为云数据库实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 指定数据库名称。例如： - oracle：serviceName.orcl。
	DbName *string `json:"db_name,omitempty"`

	// 物理源库信息。
	SourceSharding *[]BaseEndpoint `json:"source_sharding,omitempty"`
}

func (o BaseEndpoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseEndpoint struct{}"
	}

	return strings.Join([]string{"BaseEndpoint", string(data)}, " ")
}

type BaseEndpointEndpointName struct {
	value string
}

type BaseEndpointEndpointNameEnum struct {
	ORACLE              BaseEndpointEndpointName
	ECS_ORACLE          BaseEndpointEndpointName
	CLOUD_GAUSSDBV5     BaseEndpointEndpointName
	MYSQL               BaseEndpointEndpointName
	ECS_MYSQL           BaseEndpointEndpointName
	CLOUD_MYSQL         BaseEndpointEndpointName
	REDIS               BaseEndpointEndpointName
	ECS_REDIS           BaseEndpointEndpointName
	REDISCLUSTER        BaseEndpointEndpointName
	ECS_REDISCLUSTER    BaseEndpointEndpointName
	CLOUD_GAUSSDB_REDIS BaseEndpointEndpointName
	POSTGRESQL          BaseEndpointEndpointName
	ECS_POSTGRESQL      BaseEndpointEndpointName
	CLOUD_POSTGRESQL    BaseEndpointEndpointName
	MONGODB             BaseEndpointEndpointName
	ECS_MONGODB         BaseEndpointEndpointName
	CLOUD_MONGODB       BaseEndpointEndpointName
}

func GetBaseEndpointEndpointNameEnum() BaseEndpointEndpointNameEnum {
	return BaseEndpointEndpointNameEnum{
		ORACLE: BaseEndpointEndpointName{
			value: "oracle",
		},
		ECS_ORACLE: BaseEndpointEndpointName{
			value: "ecs_oracle",
		},
		CLOUD_GAUSSDBV5: BaseEndpointEndpointName{
			value: "cloud_gaussdbv5",
		},
		MYSQL: BaseEndpointEndpointName{
			value: "mysql",
		},
		ECS_MYSQL: BaseEndpointEndpointName{
			value: "ecs_mysql",
		},
		CLOUD_MYSQL: BaseEndpointEndpointName{
			value: "cloud_mysql",
		},
		REDIS: BaseEndpointEndpointName{
			value: "redis",
		},
		ECS_REDIS: BaseEndpointEndpointName{
			value: "ecs_redis",
		},
		REDISCLUSTER: BaseEndpointEndpointName{
			value: "rediscluster",
		},
		ECS_REDISCLUSTER: BaseEndpointEndpointName{
			value: "ecs_rediscluster",
		},
		CLOUD_GAUSSDB_REDIS: BaseEndpointEndpointName{
			value: "cloud_gaussdb_redis",
		},
		POSTGRESQL: BaseEndpointEndpointName{
			value: "postgresql",
		},
		ECS_POSTGRESQL: BaseEndpointEndpointName{
			value: "ecs_postgresql",
		},
		CLOUD_POSTGRESQL: BaseEndpointEndpointName{
			value: "cloud_postgresql",
		},
		MONGODB: BaseEndpointEndpointName{
			value: "mongodb",
		},
		ECS_MONGODB: BaseEndpointEndpointName{
			value: "ecs_mongodb",
		},
		CLOUD_MONGODB: BaseEndpointEndpointName{
			value: "cloud_mongodb",
		},
	}
}

func (c BaseEndpointEndpointName) Value() string {
	return c.value
}

func (c BaseEndpointEndpointName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseEndpointEndpointName) UnmarshalJSON(b []byte) error {
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
