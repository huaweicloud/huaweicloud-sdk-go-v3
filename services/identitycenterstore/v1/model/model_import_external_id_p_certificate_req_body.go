package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImportExternalIdPCertificateReqBody struct {

	// PEM格式的身份提供商证书内容
	X509CertificateInPem string `json:"x509_certificate_in_pem"`

	// 身份提供商证书用途，目前仅支持签名
	CertificateUse ImportExternalIdPCertificateReqBodyCertificateUse `json:"certificate_use"`
}

func (o ImportExternalIdPCertificateReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportExternalIdPCertificateReqBody struct{}"
	}

	return strings.Join([]string{"ImportExternalIdPCertificateReqBody", string(data)}, " ")
}

type ImportExternalIdPCertificateReqBodyCertificateUse struct {
	value string
}

type ImportExternalIdPCertificateReqBodyCertificateUseEnum struct {
	SIGNING ImportExternalIdPCertificateReqBodyCertificateUse
}

func GetImportExternalIdPCertificateReqBodyCertificateUseEnum() ImportExternalIdPCertificateReqBodyCertificateUseEnum {
	return ImportExternalIdPCertificateReqBodyCertificateUseEnum{
		SIGNING: ImportExternalIdPCertificateReqBodyCertificateUse{
			value: "SIGNING",
		},
	}
}

func (c ImportExternalIdPCertificateReqBodyCertificateUse) Value() string {
	return c.value
}

func (c ImportExternalIdPCertificateReqBodyCertificateUse) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportExternalIdPCertificateReqBodyCertificateUse) UnmarshalJSON(b []byte) error {
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
