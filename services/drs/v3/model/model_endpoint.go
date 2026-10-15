package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Endpoint 数据库信息体
type Endpoint struct {

	// **参数解释**：  数据库类型，标识源库或目标库的数据库引擎种类。  **约束限制**：  测试连接之后修改调用时必填。  **取值范围**：  - mysql：MySQL - mongodb：MongoDB - gaussdbv5：GaussDB分布式版 - postgresql：PostgreSQL - kafka：Kafka - gaussdbv5ha：GaussDB集中式版 - taurus：TaurusDB  **默认取值**：  不涉及
	DbType *EndpointDbType `json:"db_type,omitempty"`

	// **参数解释**：  数据库所在的可用区（AZ，Availability Zone）编码。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	AzCode *string `json:"az_code,omitempty"`

	// **参数解释**：  数据库实例所在区域。数据库为云上实例时，比如RDS实例时必填。  **约束限制**：  - 灾备场景下job_direction为down时source_endpoint中该值为必填。 - 灾备场景下job_direction为up时target_endpoint中该值为必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**：  数据库实例ID，用于标识华为云RDS数据库实例。数据库为云上实例时，比如RDS实例时必填。  **约束限制**：  - 灾备场景下job_direction为down时source_endpoint中该值为必填。 - 灾备场景下job_direction为up时target_endpoint中该值为必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	InstId *string `json:"inst_id,omitempty"`

	// **参数解释**：  数据库所在的虚拟私有云（VPC，Virtual Private Cloud）ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：  数据库所在的子网ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**：  数据库所在的安全组ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// **参数解释**：  数据库实例所在Region下的项目ID。如果数据库为云上实例时，填写数据库实例所在Region下的Project ID；否则，填写当前任务所在Region下的Project ID。获取方法请参见获取项目ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：  数据库名称，Oracle场景下为serviceName。  **约束限制**：  - 不能超过128位。 - 不能包含!<>&'\"特殊字符。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**：  数据库密码，用于连接数据库的认证密码。  **约束限制**：  除不需要密码的引擎，其他引擎测试连接之后修改调用时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DbPassword *string `json:"db_password,omitempty"`

	// **参数解释**：  数据库端口号，用于指定连接数据库的服务端口。  **约束限制**：  输入范围为1-65535之间的整数。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DbPort *int32 `json:"db_port,omitempty"`

	// **参数解释**：  数据库帐号，用于连接数据库的用户名。  **约束限制**：  除不需要用户名的引擎，其他引擎测试连接之后修改调用时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DbUser *string `json:"db_user,omitempty"`

	// **参数解释**：  RDS实例名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	InstName *string `json:"inst_name,omitempty"`

	// **参数解释**：  数据库IP地址。测试连接之后修改调用时必填。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Ip *string `json:"ip,omitempty"`

	// **参数解释**：  MongoDB高可用模式。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	MongoHaMode *string `json:"mongo_ha_mode,omitempty"`

	// **参数解释**：  MRS集群运行模式。  **约束限制**：  不涉及。  **取值范围**：  - 0：普通集群 - 1：安全集群  **默认取值**：  不涉及
	SafeMode *int32 `json:"safe_mode,omitempty"`

	// **参数解释**：  SSL证书密码，证书文件后缀为.p12。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SslCertPassword *string `json:"ssl_cert_password,omitempty"`

	// **参数解释**：  SSL证书内容checksum值，后端校验，源库安全连接必选。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SslCertCheckSum *string `json:"ssl_cert_check_sum,omitempty"`

	// **参数解释**：  SSL证书内容，用base64加密。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SslCertKey *string `json:"ssl_cert_key,omitempty"`

	// **参数解释**：  SSL证书名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SslCertName *string `json:"ssl_cert_name,omitempty"`

	// **参数解释**：  是否使用SSL安全连接，指定数据库连接是否启用SSL加密传输。  **约束限制**：  不涉及。  **取值范围**：  - true：使用SSL安全连接 - false：不使用SSL安全连接  **默认取值**：  不涉及
	SslLink *bool `json:"ssl_link,omitempty"`

	// **参数解释**：  Kafka Topic名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Topic *string `json:"topic,omitempty"`

	// **参数解释**：  MongoDB集群模式。当集群实例无法获取到分片节点的IP时，source_endpoint中需要填写。  **约束限制**：  不涉及。  **取值范围**：  - Sharding4.0+：MongoDB集群4.0及以上版本  **默认取值**：  Sharding4.0+
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
