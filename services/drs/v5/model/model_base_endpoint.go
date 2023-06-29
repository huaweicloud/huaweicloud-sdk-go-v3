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

	// 数据库场景类型。取值： - oracle：云下自建Oracle数据库。 - ecs_oracle：华为云ECS自建Oracle数据库。 - cloud_gaussdbv5：华为云数据库GaussDB分布式。 - mysql：他云/本地自建MySQL数据库。 - ecs_mysql：华为云ECS自建MySQL数据库。 - cloud_mysql：华为云数据库RDS for MySQL。
	EndpointName BaseEndpointEndpointName `json:"endpoint_name"`

	// 数据库IP。 约束： - 数据库为自建MongoDB时，数据库IP与端口之间用“:”英文冒号拼接，多个值之间请用“,”英文逗号隔开，最多支持填写3个IP地址或域名。 - 数据库为DDS实例时，数据库IP与端口之间用“:”英文冒号拼接，多个IP端口之间请用“,”英文逗号分隔。 示例： - MySQL：192.168.0.10 - MongoDB：192.168.0.10:8080,192.168.0.11:8080,192.168.0.12:8080 - DDS：192.168.205.130:8635,192.168.250.64:8635
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
	ORACLE          BaseEndpointEndpointName
	ECS_ORACLE      BaseEndpointEndpointName
	CLOUD_GAUSSDBV5 BaseEndpointEndpointName
	MYSQL           BaseEndpointEndpointName
	ECS_MYSQL       BaseEndpointEndpointName
	CLOUD_MYSQL     BaseEndpointEndpointName
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
