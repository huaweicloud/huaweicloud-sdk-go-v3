package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobEndpointInfo 创建任务数据库信息体。
type JobEndpointInfo struct {

	// **参数解释**：  数据库类型，标识源库或目标库的数据库引擎种类。  **约束限制**：  不涉及。  **取值范围**：  - oracle：Oracle - gaussdbv5：GaussDB分布式版 - redis：Redis - rediscluster：Redis集群版 - gaussredis：GeminiDB Redis - mysql：MySQL  **默认取值**：  不涉及
	DbType JobEndpointInfoDbType `json:"db_type"`

	// **参数解释**：  数据库实例类型，标识数据库的部署方式。  **约束限制**：  不涉及。  **取值范围**：  - offline：自建数据库 - ecs：华为云ECS自建数据库 - cloud：华为云数据库  **默认取值**：  不涉及
	EndpointType JobEndpointInfoEndpointType `json:"endpoint_type"`

	// **参数解释**：  数据库实例角色，标识该数据库在数据复制任务中是源端还是目标端。  **约束限制**：  不涉及。  **取值范围**：  - so：源库 - ta：目标库  **默认取值**：  不涉及
	EndpointRole JobEndpointInfoEndpointRole `json:"endpoint_role"`

	Endpoint *BaseEndpoint `json:"endpoint"`

	Cloud *CloudBaseInfo `json:"cloud,omitempty"`

	Vpc *CloudVpcInfo `json:"vpc,omitempty"`

	Config *BaseEndpointConfig `json:"config,omitempty"`

	Ssl *EndpointSslConfig `json:"ssl,omitempty"`

	CustomizedDns *CustomizedDns `json:"customized_dns,omitempty"`
}

func (o JobEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEndpointInfo struct{}"
	}

	return strings.Join([]string{"JobEndpointInfo", string(data)}, " ")
}

type JobEndpointInfoDbType struct {
	value string
}

type JobEndpointInfoDbTypeEnum struct {
	ORACLE       JobEndpointInfoDbType
	GAUSSDBV5    JobEndpointInfoDbType
	REDIS        JobEndpointInfoDbType
	REDISCLUSTER JobEndpointInfoDbType
	GAUSSREDIS   JobEndpointInfoDbType
	MYSQL        JobEndpointInfoDbType
}

func GetJobEndpointInfoDbTypeEnum() JobEndpointInfoDbTypeEnum {
	return JobEndpointInfoDbTypeEnum{
		ORACLE: JobEndpointInfoDbType{
			value: "oracle",
		},
		GAUSSDBV5: JobEndpointInfoDbType{
			value: "gaussdbv5",
		},
		REDIS: JobEndpointInfoDbType{
			value: "redis",
		},
		REDISCLUSTER: JobEndpointInfoDbType{
			value: "rediscluster",
		},
		GAUSSREDIS: JobEndpointInfoDbType{
			value: "gaussredis",
		},
		MYSQL: JobEndpointInfoDbType{
			value: "mysql",
		},
	}
}

func (c JobEndpointInfoDbType) Value() string {
	return c.value
}

func (c JobEndpointInfoDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobEndpointInfoDbType) UnmarshalJSON(b []byte) error {
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

type JobEndpointInfoEndpointType struct {
	value string
}

type JobEndpointInfoEndpointTypeEnum struct {
	OFFLINE JobEndpointInfoEndpointType
	ECS     JobEndpointInfoEndpointType
	CLOUD   JobEndpointInfoEndpointType
}

func GetJobEndpointInfoEndpointTypeEnum() JobEndpointInfoEndpointTypeEnum {
	return JobEndpointInfoEndpointTypeEnum{
		OFFLINE: JobEndpointInfoEndpointType{
			value: "offline",
		},
		ECS: JobEndpointInfoEndpointType{
			value: "ecs",
		},
		CLOUD: JobEndpointInfoEndpointType{
			value: "cloud",
		},
	}
}

func (c JobEndpointInfoEndpointType) Value() string {
	return c.value
}

func (c JobEndpointInfoEndpointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobEndpointInfoEndpointType) UnmarshalJSON(b []byte) error {
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

type JobEndpointInfoEndpointRole struct {
	value string
}

type JobEndpointInfoEndpointRoleEnum struct {
	SO JobEndpointInfoEndpointRole
	TA JobEndpointInfoEndpointRole
}

func GetJobEndpointInfoEndpointRoleEnum() JobEndpointInfoEndpointRoleEnum {
	return JobEndpointInfoEndpointRoleEnum{
		SO: JobEndpointInfoEndpointRole{
			value: "so",
		},
		TA: JobEndpointInfoEndpointRole{
			value: "ta",
		},
	}
}

func (c JobEndpointInfoEndpointRole) Value() string {
	return c.value
}

func (c JobEndpointInfoEndpointRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobEndpointInfoEndpointRole) UnmarshalJSON(b []byte) error {
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
