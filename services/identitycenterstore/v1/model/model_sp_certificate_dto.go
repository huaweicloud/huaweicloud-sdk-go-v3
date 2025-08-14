package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SpCertificateDto 服务提供商证书详情
type SpCertificateDto struct {

	// 证书ID
	CertificateId string `json:"certificate_id"`

	// x509证书
	X509certificate string `json:"x509certificate"`

	// 签名算法
	Algorithm string `json:"algorithm"`

	// 证书过期时间戳
	ExpiryDate int64 `json:"expiry_date"`

	// 证书激活状态
	State SpCertificateDtoState `json:"state"`
}

func (o SpCertificateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpCertificateDto struct{}"
	}

	return strings.Join([]string{"SpCertificateDto", string(data)}, " ")
}

type SpCertificateDtoState struct {
	value string
}

type SpCertificateDtoStateEnum struct {
	ACTIVE   SpCertificateDtoState
	INACTIVE SpCertificateDtoState
	EXPIRED  SpCertificateDtoState
}

func GetSpCertificateDtoStateEnum() SpCertificateDtoStateEnum {
	return SpCertificateDtoStateEnum{
		ACTIVE: SpCertificateDtoState{
			value: "ACTIVE",
		},
		INACTIVE: SpCertificateDtoState{
			value: "INACTIVE",
		},
		EXPIRED: SpCertificateDtoState{
			value: "EXPIRED",
		},
	}
}

func (c SpCertificateDtoState) Value() string {
	return c.value
}

func (c SpCertificateDtoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SpCertificateDtoState) UnmarshalJSON(b []byte) error {
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
