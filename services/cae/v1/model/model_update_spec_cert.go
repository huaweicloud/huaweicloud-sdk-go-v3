package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateSpecCert struct {

	// 证书内容。
	Crt string `json:"crt"`

	// 私钥内容。
	Key string `json:"key"`

	// 安全策略。 - tls-1-2-strict - tls-1-2 - tls-1-1 - tls-1-0
	Policy UpdateSpecCertPolicy `json:"policy"`
}

func (o UpdateSpecCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSpecCert struct{}"
	}

	return strings.Join([]string{"UpdateSpecCert", string(data)}, " ")
}

type UpdateSpecCertPolicy struct {
	value string
}

type UpdateSpecCertPolicyEnum struct {
	TLS_1_2_STRICT UpdateSpecCertPolicy
	TLS_1_2        UpdateSpecCertPolicy
	TLS_1_1        UpdateSpecCertPolicy
	TLS_1_0        UpdateSpecCertPolicy
}

func GetUpdateSpecCertPolicyEnum() UpdateSpecCertPolicyEnum {
	return UpdateSpecCertPolicyEnum{
		TLS_1_2_STRICT: UpdateSpecCertPolicy{
			value: "tls-1-2-strict",
		},
		TLS_1_2: UpdateSpecCertPolicy{
			value: "tls-1-2",
		},
		TLS_1_1: UpdateSpecCertPolicy{
			value: "tls-1-1",
		},
		TLS_1_0: UpdateSpecCertPolicy{
			value: "tls-1-0",
		},
	}
}

func (c UpdateSpecCertPolicy) Value() string {
	return c.value
}

func (c UpdateSpecCertPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSpecCertPolicy) UnmarshalJSON(b []byte) error {
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
