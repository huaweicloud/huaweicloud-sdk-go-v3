package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateConnectionReq 创建连接请求体。
type CreateConnectionReq struct {

	// 连接名称。 约束：连接名称在4位到50位之间，不区分大小写，可以包含字母、数字、中划线或下划线，不能包括其他特殊字符。
	Name string `json:"name"`

	// 连接类型，包含： - mysql - postgresql - mongodb - oracle
	DbType CreateConnectionReqDbType `json:"db_type"`

	Config *ConnectionConfig `json:"config,omitempty"`

	// 描述，长度不能超过255个字符。
	Description *string `json:"description,omitempty"`

	Endpoint *BaseEndpoint `json:"endpoint"`

	Vpc *CloudVpcInfo `json:"vpc,omitempty"`

	Ssl *EndpointSslConfig `json:"ssl,omitempty"`

	Cloud *CloudBaseInfo `json:"cloud,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionReq struct{}"
	}

	return strings.Join([]string{"CreateConnectionReq", string(data)}, " ")
}

type CreateConnectionReqDbType struct {
	value string
}

type CreateConnectionReqDbTypeEnum struct {
	MYSQL      CreateConnectionReqDbType
	POSTGRESQL CreateConnectionReqDbType
	MONGODB    CreateConnectionReqDbType
	ORACLE     CreateConnectionReqDbType
}

func GetCreateConnectionReqDbTypeEnum() CreateConnectionReqDbTypeEnum {
	return CreateConnectionReqDbTypeEnum{
		MYSQL: CreateConnectionReqDbType{
			value: "mysql",
		},
		POSTGRESQL: CreateConnectionReqDbType{
			value: "postgresql",
		},
		MONGODB: CreateConnectionReqDbType{
			value: "mongodb",
		},
		ORACLE: CreateConnectionReqDbType{
			value: "oracle",
		},
	}
}

func (c CreateConnectionReqDbType) Value() string {
	return c.value
}

func (c CreateConnectionReqDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateConnectionReqDbType) UnmarshalJSON(b []byte) error {
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
