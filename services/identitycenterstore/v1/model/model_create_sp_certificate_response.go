package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSpCertificateResponse Response Object
type CreateSpCertificateResponse struct {

	// 证书ID
	CertificateId *string `json:"certificate_id,omitempty"`

	// x509证书
	X509certificate *string `json:"x509certificate,omitempty"`

	// 签名算法
	Algorithm *string `json:"algorithm,omitempty"`

	// 证书过期时间戳
	ExpiryDate *int64 `json:"expiry_date,omitempty"`

	// 证书激活状态
	State          *CreateSpCertificateResponseState `json:"state,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o CreateSpCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSpCertificateResponse struct{}"
	}

	return strings.Join([]string{"CreateSpCertificateResponse", string(data)}, " ")
}

type CreateSpCertificateResponseState struct {
	value string
}

type CreateSpCertificateResponseStateEnum struct {
	ACTIVE   CreateSpCertificateResponseState
	INACTIVE CreateSpCertificateResponseState
	EXPIRED  CreateSpCertificateResponseState
}

func GetCreateSpCertificateResponseStateEnum() CreateSpCertificateResponseStateEnum {
	return CreateSpCertificateResponseStateEnum{
		ACTIVE: CreateSpCertificateResponseState{
			value: "ACTIVE",
		},
		INACTIVE: CreateSpCertificateResponseState{
			value: "INACTIVE",
		},
		EXPIRED: CreateSpCertificateResponseState{
			value: "EXPIRED",
		},
	}
}

func (c CreateSpCertificateResponseState) Value() string {
	return c.value
}

func (c CreateSpCertificateResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSpCertificateResponseState) UnmarshalJSON(b []byte) error {
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
