package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CertificateForm 添加或编辑证书的请求体表单
type CertificateForm struct {

	// 证书名称。支持中文，英文字母，数字，下划线，且只能以英文或汉字开头，4~50个字符。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// 证书内容
	CertContent string `json:"cert_content"`

	// 证书私钥
	PrivateKey string `json:"private_key"`

	// 证书类型。 - instance：实例证书 - global：全局证书
	Type *CertificateFormType `json:"type,omitempty"`

	// 所属实例ID，当type=instance时必填
	InstanceId *string `json:"instance_id,omitempty"`

	// 信任的根证书CA
	TrustedRootCa *string `json:"trusted_root_ca,omitempty"`

	// 证书算法类型： - RSA - ECC - SM2
	AlgorithmType *CertificateFormAlgorithmType `json:"algorithm_type,omitempty"`

	// 签名类型证书内容，仅algorithm_type=SM2时必填。
	CertContentSign *string `json:"cert_content_sign,omitempty"`

	// 签名类型私钥内容，仅algorithm_type=SM2时必填。
	PrivateKeySign *string `json:"private_key_sign,omitempty"`
}

func (o CertificateForm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateForm struct{}"
	}

	return strings.Join([]string{"CertificateForm", string(data)}, " ")
}

type CertificateFormType struct {
	value string
}

type CertificateFormTypeEnum struct {
	INSTANCE CertificateFormType
	GLOBAL   CertificateFormType
}

func GetCertificateFormTypeEnum() CertificateFormTypeEnum {
	return CertificateFormTypeEnum{
		INSTANCE: CertificateFormType{
			value: "instance",
		},
		GLOBAL: CertificateFormType{
			value: "global",
		},
	}
}

func (c CertificateFormType) Value() string {
	return c.value
}

func (c CertificateFormType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertificateFormType) UnmarshalJSON(b []byte) error {
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

type CertificateFormAlgorithmType struct {
	value string
}

type CertificateFormAlgorithmTypeEnum struct {
	RSA CertificateFormAlgorithmType
	ECC CertificateFormAlgorithmType
	SM2 CertificateFormAlgorithmType
}

func GetCertificateFormAlgorithmTypeEnum() CertificateFormAlgorithmTypeEnum {
	return CertificateFormAlgorithmTypeEnum{
		RSA: CertificateFormAlgorithmType{
			value: "RSA",
		},
		ECC: CertificateFormAlgorithmType{
			value: "ECC",
		},
		SM2: CertificateFormAlgorithmType{
			value: "SM2",
		},
	}
}

func (c CertificateFormAlgorithmType) Value() string {
	return c.value
}

func (c CertificateFormAlgorithmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertificateFormAlgorithmType) UnmarshalJSON(b []byte) error {
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
