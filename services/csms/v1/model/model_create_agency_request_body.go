package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAgencyRequestBody 创建服务委托的请求消息体。
type CreateAgencyRequestBody struct {

	// 凭据类型。
	SecretType CreateAgencyRequestBodySecretType `json:"secret_type"`
}

func (o CreateAgencyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAgencyRequestBody", string(data)}, " ")
}

type CreateAgencyRequestBodySecretType struct {
	value string
}

type CreateAgencyRequestBodySecretTypeEnum struct {
	RDS_FG      CreateAgencyRequestBodySecretType
	GAUSS_DB_FG CreateAgencyRequestBodySecretType
}

func GetCreateAgencyRequestBodySecretTypeEnum() CreateAgencyRequestBodySecretTypeEnum {
	return CreateAgencyRequestBodySecretTypeEnum{
		RDS_FG: CreateAgencyRequestBodySecretType{
			value: "RDS-FG",
		},
		GAUSS_DB_FG: CreateAgencyRequestBodySecretType{
			value: "GaussDB-FG",
		},
	}
}

func (c CreateAgencyRequestBodySecretType) Value() string {
	return c.value
}

func (c CreateAgencyRequestBodySecretType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAgencyRequestBodySecretType) UnmarshalJSON(b []byte) error {
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
