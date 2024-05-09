package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Endpoint 数据库信息体
type Endpoint struct {

	// 数据库类型，测试连接之后修改调用时必填。
	DbType *EndpointDbType `json:"db_type,omitempty"`

	// 数据库所在可用区azCode
	AzCode *string `json:"az_code,omitempty"`

	// RDS实例所在Region，数据库为RDS实例时必填（灾备场景下job_direction为down时source_endpoint中该值为必填，job_direction为up时target_endpoint中该值为必填）。
	Region *string `json:"region,omitempty"`

	// RDS实例ID，数据库为RDS实例必填（灾备场景下job_direction为down时source_endpoint中该值为必填，job_direction为up时target_endpoint中该值为必填）。
	InstId *string `json:"inst_id,omitempty"`

	// 数据库所在的虚拟私有云id
	VpcId *string `json:"vpc_id,omitempty"`

	// 数据库所在的子网id
	SubnetId *string `json:"subnet_id,omitempty"`

	// 数据库所在的安全组id。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// RDS实例projectId
	ProjectId *string `json:"project_id,omitempty"`

	// 服务名serviceName，源库为oracle场景时必填。约束：不能超过128位，不能包含!<>&'\"\\特殊字符。待还原数据库名称是指备份文件中包含的数据库名称，当您选择部分数据库恢复时，需要选择恢复一个或者多个数据库。
	DbName *string `json:"db_name,omitempty"`

	// 数据库密码。
	DbPassword *string `json:"db_password,omitempty"`

	// 数据库端口。约束：输入范围为1-65535之间的整数。
	DbPort *int32 `json:"db_port,omitempty"`

	// 数据库用户。
	DbUser *string `json:"db_user,omitempty"`

	// RDS实例名称。
	InstName *string `json:"inst_name,omitempty"`

	// 数据库ip
	Ip *string `json:"ip,omitempty"`

	// mongo ha模式。
	MongoHaMode *string `json:"mongo_ha_mode,omitempty"`

	// MRS集群运行模式，取值： - 0普通集群 - 1安全集群
	SafeMode *int32 `json:"safe_mode,omitempty"`

	// SSL证书密码，证书文件后缀为.p12
	SslCertPassword *string `json:"ssl_cert_password,omitempty"`

	// SSL证书内容checksum值，后端校验，源库安全连接必选。
	SslCertCheckSum *string `json:"ssl_cert_check_sum,omitempty"`

	// SSL证书内容，用base64加密
	SslCertKey *string `json:"ssl_cert_key,omitempty"`

	// SSL证书名字
	SslCertName *string `json:"ssl_cert_name,omitempty"`

	// 是否SSL安全连接。
	SslLink *bool `json:"ssl_link,omitempty"`

	// kafka topic名称
	Topic *string `json:"topic,omitempty"`

	// MongDB集群4.0及以上版本，当集群实例无法获取到分片节点的IP时，source_endpoint中需要填写，值为：Sharding4.0+。
	ClusterMode *EndpointClusterMode `json:"cluster_mode,omitempty"`

	KafkaSecurityConfig *KafkaSecurity `json:"kafka_security_config,omitempty"`
}

func (o Endpoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Endpoint struct{}"
	}

	return strings.Join([]string{"Endpoint", string(data)}, " ")
}

type EndpointDbType struct {
	value string
}

type EndpointDbTypeEnum struct {
	MYSQL       EndpointDbType
	MONGODB     EndpointDbType
	GAUSSDBV5   EndpointDbType
	POSTGRESQL  EndpointDbType
	KAFKA       EndpointDbType
	GAUSSDBV5HA EndpointDbType
	TAURUS      EndpointDbType
}

func GetEndpointDbTypeEnum() EndpointDbTypeEnum {
	return EndpointDbTypeEnum{
		MYSQL: EndpointDbType{
			value: "mysql",
		},
		MONGODB: EndpointDbType{
			value: "mongodb",
		},
		GAUSSDBV5: EndpointDbType{
			value: "gaussdbv5",
		},
		POSTGRESQL: EndpointDbType{
			value: "postgresql",
		},
		KAFKA: EndpointDbType{
			value: "kafka",
		},
		GAUSSDBV5HA: EndpointDbType{
			value: "gaussdbv5ha",
		},
		TAURUS: EndpointDbType{
			value: "taurus",
		},
	}
}

func (c EndpointDbType) Value() string {
	return c.value
}

func (c EndpointDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointDbType) UnmarshalJSON(b []byte) error {
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

type EndpointClusterMode struct {
	value string
}

type EndpointClusterModeEnum struct {
	SHARDING4_0 EndpointClusterMode
}

func GetEndpointClusterModeEnum() EndpointClusterModeEnum {
	return EndpointClusterModeEnum{
		SHARDING4_0: EndpointClusterMode{
			value: "Sharding4.0+",
		},
	}
}

func (c EndpointClusterMode) Value() string {
	return c.value
}

func (c EndpointClusterMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointClusterMode) UnmarshalJSON(b []byte) error {
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
