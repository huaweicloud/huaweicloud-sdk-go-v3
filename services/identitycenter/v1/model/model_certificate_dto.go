package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CertificateDto 应用程序证书
type CertificateDto struct {

	// 证书生成算法
	Algorithm string `json:"algorithm"`

	// 应用程序证书
	Certificate string `json:"certificate"`

	// 应用程序证书Id
	CertificateId string `json:"certificate_id"`

	// 证书过期时间
	ExpiryDate int64 `json:"expiry_date"`

	// 证书状态
	Status CertificateDtoStatus `json:"status"`

	// 密钥大小
	KeySize string `json:"key_size"`

	// 证书生成时间
	IssueDate int64 `json:"issue_date"`
}

func (o CertificateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateDto struct{}"
	}

	return strings.Join([]string{"CertificateDto", string(data)}, " ")
}

type CertificateDtoStatus struct {
	value string
}

type CertificateDtoStatusEnum struct {
	ACTIVE   CertificateDtoStatus
	INACTIVE CertificateDtoStatus
}

func GetCertificateDtoStatusEnum() CertificateDtoStatusEnum {
	return CertificateDtoStatusEnum{
		ACTIVE: CertificateDtoStatus{
			value: "ACTIVE",
		},
		INACTIVE: CertificateDtoStatus{
			value: "INACTIVE",
		},
	}
}

func (c CertificateDtoStatus) Value() string {
	return c.value
}

func (c CertificateDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertificateDtoStatus) UnmarshalJSON(b []byte) error {
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
