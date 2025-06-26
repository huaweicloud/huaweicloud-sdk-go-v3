package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateSignaturePolicyRequestBody struct {

	// 签名策略名称，由字母、汉字、数字、下划线（_）、中划线 (-)组成，1-256个字符。
	Name string `json:"name"`

	// 签名策略描述
	Description *string `json:"description,omitempty"`

	// 是否开启
	Enabled bool `json:"enabled"`

	// 加签方式，可选KMS
	SignatureMethod CreateSignaturePolicyRequestBodySignatureMethod `json:"signature_method"`

	// 加签算法，KMS的密钥算法EC_P256对应着ECDSA_SHA_256，EC_P384对应着ECDSA_SHA_384，SM2对应着SM2DSA_SM3
	SignatureAlgorithm CreateSignaturePolicyRequestBodySignatureAlgorithm `json:"signature_algorithm"`

	// 加签Key
	SignatureKey string `json:"signature_key"`

	Trigger *TriggerConfig `json:"trigger"`

	// 作用范围规则
	ScopeRules []SignScopeRule `json:"scope_rules"`
}

func (o CreateSignaturePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSignaturePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSignaturePolicyRequestBody", string(data)}, " ")
}

type CreateSignaturePolicyRequestBodySignatureMethod struct {
	value string
}

type CreateSignaturePolicyRequestBodySignatureMethodEnum struct {
	KMS CreateSignaturePolicyRequestBodySignatureMethod
}

func GetCreateSignaturePolicyRequestBodySignatureMethodEnum() CreateSignaturePolicyRequestBodySignatureMethodEnum {
	return CreateSignaturePolicyRequestBodySignatureMethodEnum{
		KMS: CreateSignaturePolicyRequestBodySignatureMethod{
			value: "KMS",
		},
	}
}

func (c CreateSignaturePolicyRequestBodySignatureMethod) Value() string {
	return c.value
}

func (c CreateSignaturePolicyRequestBodySignatureMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSignaturePolicyRequestBodySignatureMethod) UnmarshalJSON(b []byte) error {
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

type CreateSignaturePolicyRequestBodySignatureAlgorithm struct {
	value string
}

type CreateSignaturePolicyRequestBodySignatureAlgorithmEnum struct {
	ECDSA_SHA_384 CreateSignaturePolicyRequestBodySignatureAlgorithm
	SM2_DSA_SM3   CreateSignaturePolicyRequestBodySignatureAlgorithm
	ECDSA_SHA_256 CreateSignaturePolicyRequestBodySignatureAlgorithm
}

func GetCreateSignaturePolicyRequestBodySignatureAlgorithmEnum() CreateSignaturePolicyRequestBodySignatureAlgorithmEnum {
	return CreateSignaturePolicyRequestBodySignatureAlgorithmEnum{
		ECDSA_SHA_384: CreateSignaturePolicyRequestBodySignatureAlgorithm{
			value: "ECDSA_SHA_384",
		},
		SM2_DSA_SM3: CreateSignaturePolicyRequestBodySignatureAlgorithm{
			value: "SM2DSA_SM3",
		},
		ECDSA_SHA_256: CreateSignaturePolicyRequestBodySignatureAlgorithm{
			value: "ECDSA_SHA_256",
		},
	}
}

func (c CreateSignaturePolicyRequestBodySignatureAlgorithm) Value() string {
	return c.value
}

func (c CreateSignaturePolicyRequestBodySignatureAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSignaturePolicyRequestBodySignatureAlgorithm) UnmarshalJSON(b []byte) error {
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
