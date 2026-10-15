package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TestEndPoint 测试连接信息体
type TestEndPoint struct {

	// **参数解释**：  DRS任务ID，用于唯一标识一个数据复制任务，可从任务列表或任务详情页面获取。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Id string `json:"id"`

	// **参数解释**：  任务使用的网络类型。  **约束限制**：  不涉及。  **取值范围**：  - vpn：虚拟专用网络 - vpc：虚拟私有云 - eip：弹性公网IP  **默认取值**：  不涉及
	NetType TestEndPointNetType `json:"net_type"`

	// **参数解释**：  数据库类型，标识源库或目标库的数据库引擎种类。  **约束限制**：  不涉及。  **取值范围**：  - mysql：MySQL - mongodb：MongoDB - gaussdbv5：GaussDB分布式版 - postgresql：PostgreSQL - kafka：Kafka - gaussdbv5ha：GaussDB集中式版 - taurus：TaurusDB  **默认取值**：  不涉及
	DbType TestEndPointDbType `json:"db_type"`

	// **参数解释**：  数据库IP地址，用于指定连接数据库的服务器地址。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Ip string `json:"ip"`

	// **参数解释**：  数据库端口号，用于指定连接数据库的服务端口。  **约束限制**：  - db_type为MongoDB、DDS时端口必填取值为0。 - db_type为Oracle时为必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DbPort *int32 `json:"db_port,omitempty"`

	// **参数解释**：  RDS实例ID，用于标识华为云RDS数据库实例。  **约束限制**：  - 数据库为RDS实例时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	InstId *string `json:"inst_id,omitempty"`

	// **参数解释**：  数据库帐号，用于连接数据库的用户名。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DbUser string `json:"db_user"`

	// **参数解释**：  数据库密码，用于连接数据库的认证密码。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DbPassword string `json:"db_password"`

	// **参数解释**：  是否使用SSL安全连接，指定数据库连接是否启用SSL加密传输。  **约束限制**：  - 该参数为true时，需填写SSL证书相关信息参数。  **取值范围**：  - true：使用SSL安全连接 - false：不使用SSL安全连接  **默认取值**：  不涉及
	SslLink *bool `json:"ssl_link,omitempty"`

	// **参数解释**：  SSL证书内容，获取SSL证书后将内容通过BASE64加密后所得的字符串。  **约束限制**：  - ssl_link为true时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SslCertKey *string `json:"ssl_cert_key,omitempty"`

	// **参数解释**：  SSL证书名字。  **约束限制**：  - ssl_link为true时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SslCertName *string `json:"ssl_cert_name,omitempty"`

	// **参数解释**：  SSL证书内容checksum值，获取SSL证书后将内容经过SHA256加密后的值，后端校验用。  **约束限制**：  - ssl_link为true时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SslCertCheckSum *string `json:"ssl_cert_check_sum,omitempty"`

	// **参数解释**：  SSL证书密码。  **约束限制**：  - 证书文件后缀为.p12时需要填写密码。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SslCertPassword *string `json:"ssl_cert_password,omitempty"`

	// **参数解释**：  虚拟私有云ID，用于标识数据库所在的VPC。  **约束限制**：  - 网络类型为VPC时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：  子网ID，用于标识数据库所在的子网。  **约束限制**：  - 网络类型为VPC时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**：  数据库端点类型，标识该数据库在数据复制任务中是源端还是目标端。  **约束限制**：  不涉及。  **取值范围**：  - so：源库 - ta：目标库  **默认取值**：  so
	EndPointType TestEndPointEndPointType `json:"end_point_type"`

	// **参数解释**：  RDS实例所在区域，用于标识数据库实例的地理区域。  **约束限制**：  - 数据库为RDS实例时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**：  数据库实例所在Region的项目ID，数据库为云上实例时填写数据库实例所在Region下的项目ID，否则填写当前任务所在Region下的项目ID。获取方法请参见获取项目ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：  数据库名称，Oracle场景下为serviceName。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	DbName *string `json:"db_name,omitempty"`

	KafkaSecurityConfig *KafkaSecurity `json:"kafka_security_config,omitempty"`

	CustomizedDns *CustomizedDns `json:"customized_dns,omitempty"`
}

func (o TestEndPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestEndPoint struct{}"
	}

	return strings.Join([]string{"TestEndPoint", string(data)}, " ")
}

type TestEndPointNetType struct {
	value string
}

type TestEndPointNetTypeEnum struct {
	VPN TestEndPointNetType
	VPC TestEndPointNetType
	EIP TestEndPointNetType
}

func GetTestEndPointNetTypeEnum() TestEndPointNetTypeEnum {
	return TestEndPointNetTypeEnum{
		VPN: TestEndPointNetType{
			value: "vpn",
		},
		VPC: TestEndPointNetType{
			value: "vpc",
		},
		EIP: TestEndPointNetType{
			value: "eip",
		},
	}
}

func (c TestEndPointNetType) Value() string {
	return c.value
}

func (c TestEndPointNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TestEndPointNetType) UnmarshalJSON(b []byte) error {
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

type TestEndPointDbType struct {
	value string
}

type TestEndPointDbTypeEnum struct {
	MYSQL       TestEndPointDbType
	MONGODB     TestEndPointDbType
	GAUSSDBV5   TestEndPointDbType
	POSTGRESQL  TestEndPointDbType
	KAFKA       TestEndPointDbType
	GAUSSDBV5HA TestEndPointDbType
	TAURUS      TestEndPointDbType
}

func GetTestEndPointDbTypeEnum() TestEndPointDbTypeEnum {
	return TestEndPointDbTypeEnum{
		MYSQL: TestEndPointDbType{
			value: "mysql",
		},
		MONGODB: TestEndPointDbType{
			value: "mongodb",
		},
		GAUSSDBV5: TestEndPointDbType{
			value: "gaussdbv5",
		},
		POSTGRESQL: TestEndPointDbType{
			value: "postgresql",
		},
		KAFKA: TestEndPointDbType{
			value: "kafka",
		},
		GAUSSDBV5HA: TestEndPointDbType{
			value: "gaussdbv5ha",
		},
		TAURUS: TestEndPointDbType{
			value: "taurus",
		},
	}
}

func (c TestEndPointDbType) Value() string {
	return c.value
}

func (c TestEndPointDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TestEndPointDbType) UnmarshalJSON(b []byte) error {
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

type TestEndPointEndPointType struct {
	value string
}

type TestEndPointEndPointTypeEnum struct {
	SO TestEndPointEndPointType
	TA TestEndPointEndPointType
}

func GetTestEndPointEndPointTypeEnum() TestEndPointEndPointTypeEnum {
	return TestEndPointEndPointTypeEnum{
		SO: TestEndPointEndPointType{
			value: "so",
		},
		TA: TestEndPointEndPointType{
			value: "ta",
		},
	}
}

func (c TestEndPointEndPointType) Value() string {
	return c.value
}

func (c TestEndPointEndPointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TestEndPointEndPointType) UnmarshalJSON(b []byte) error {
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
