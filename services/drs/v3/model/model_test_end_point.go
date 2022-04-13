package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 测试连接信息体
type TestEndPoint struct {
	// 任务ID

	Id string `json:"id"`
	// 网络类型

	NetType TestEndPointNetType `json:"net_type"`
	// 数据库类型

	DbType TestEndPointDbType `json:"db_type"`
	// 数据库IP

	Ip string `json:"ip"`
	// 数据库端口，Mongo、DDS必填为0。

	DbPort *int32 `json:"db_port,omitempty"`
	// RDS实例id，RDS实例必填。

	InstId *string `json:"inst_id,omitempty"`
	// 数据库帐号。

	DbUser string `json:"db_user"`
	// 数据库密码。

	DbPassword string `json:"db_password"`
	// 是否SSL安全连接。

	SslLink *bool `json:"ssl_link,omitempty"`
	// SSL证书内容，base64加密后的值，源库安全连接必选。

	SslCertKey *string `json:"ssl_cert_key,omitempty"`
	// SSL证书名字，源库安全连接必选。

	SslCertName *string `json:"ssl_cert_name,omitempty"`
	// SSL证书内容checksum值，证书经过sha256加密后的值，后端校验，源库安全连接必选。

	SslCertCheckSum *string `json:"ssl_cert_check_sum,omitempty"`
	// SSL证书密码，证书文件后缀为.p12，需要密码。

	SslCertPassword *string `json:"ssl_cert_password,omitempty"`
	// vpcid，数据库为RDS时必选。

	VpcId *string `json:"vpc_id,omitempty"`
	// subnetid，数据库为RDS必选。

	SubnetId *string `json:"subnet_id,omitempty"`
	// 源库：so,目标库：ta

	EndPointType TestEndPointEndPointType `json:"end_point_type"`
	// rds实例region，数据库为RDS时必填。

	Region *string `json:"region,omitempty"`
	// 用户所处region的projectId。

	ProjectId *string `json:"project_id,omitempty"`
	// 数据库用户名，DDS的账号认证数据库，Oracle的serviceName。

	DbName *string `json:"db_name,omitempty"`

	KafkaSecurityConfig *KafkaSecurity `json:"kafka_security_config,omitempty"`
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

func (c TestEndPointNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TestEndPointNetType) UnmarshalJSON(b []byte) error {
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

type TestEndPointDbType struct {
	value string
}

type TestEndPointDbTypeEnum struct {
	MYSQL      TestEndPointDbType
	MONGODB    TestEndPointDbType
	POSTGRESQL TestEndPointDbType
}

func GetTestEndPointDbTypeEnum() TestEndPointDbTypeEnum {
	return TestEndPointDbTypeEnum{
		MYSQL: TestEndPointDbType{
			value: "mysql",
		},
		MONGODB: TestEndPointDbType{
			value: "mongodb",
		},
		POSTGRESQL: TestEndPointDbType{
			value: "postgresql",
		},
	}
}

func (c TestEndPointDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TestEndPointDbType) UnmarshalJSON(b []byte) error {
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

func (c TestEndPointEndPointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TestEndPointEndPointType) UnmarshalJSON(b []byte) error {
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
