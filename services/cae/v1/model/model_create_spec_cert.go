package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateSpecCert struct {

	// 证书内容。
	Crt string `json:"crt"`

	// 私钥内容。
	Key string `json:"key"`

	// 安全策略。 - tls-1-2-strict - tls-1-2 - tls-1-1 - tls-1-0
	Policy CreateSpecCertPolicy `json:"policy"`
}

func (o CreateSpecCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSpecCert struct{}"
	}

	return strings.Join([]string{"CreateSpecCert", string(data)}, " ")
}

type CreateSpecCertPolicy struct {
	value string
}

type CreateSpecCertPolicyEnum struct {
	TLS_1_2_STRICT CreateSpecCertPolicy
	TLS_1_2        CreateSpecCertPolicy
	TLS_1_1        CreateSpecCertPolicy
	TLS_1_0        CreateSpecCertPolicy
}

func GetCreateSpecCertPolicyEnum() CreateSpecCertPolicyEnum {
	return CreateSpecCertPolicyEnum{
		TLS_1_2_STRICT: CreateSpecCertPolicy{
			value: "tls-1-2-strict",
		},
		TLS_1_2: CreateSpecCertPolicy{
			value: "tls-1-2",
		},
		TLS_1_1: CreateSpecCertPolicy{
			value: "tls-1-1",
		},
		TLS_1_0: CreateSpecCertPolicy{
			value: "tls-1-0",
		},
	}
}

func (c CreateSpecCertPolicy) Value() string {
	return c.value
}

func (c CreateSpecCertPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSpecCertPolicy) UnmarshalJSON(b []byte) error {
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
