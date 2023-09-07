package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CertificateKindObj API类型，固定值“Certificate”，该值不可修改。
type CertificateKindObj struct {
	value string
}

type CertificateKindObjEnum struct {
	CERTIFICATE CertificateKindObj
}

func GetCertificateKindObjEnum() CertificateKindObjEnum {
	return CertificateKindObjEnum{
		CERTIFICATE: CertificateKindObj{
			value: "Certificate",
		},
	}
}

func (c CertificateKindObj) Value() string {
	return c.value
}

func (c CertificateKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertificateKindObj) UnmarshalJSON(b []byte) error {
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
