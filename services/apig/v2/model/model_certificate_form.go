package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CertificateForm 添加或编辑证书的请求体表单
type CertificateForm struct {

	// 证书名称
	Name string `json:"name"`

	// 证书内容
	CertContent string `json:"cert_content"`

	// 证书私钥
	PrivateKey string `json:"private_key"`

	// 证书可见范围
	Type *CertificateFormType `json:"type,omitempty"`

	// 所属实例ID，当type=instance时必填
	InstanceId *string `json:"instance_id,omitempty"`

	// 信任的根证书CA
	TrustedRootCa *string `json:"trusted_root_ca,omitempty"`
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
