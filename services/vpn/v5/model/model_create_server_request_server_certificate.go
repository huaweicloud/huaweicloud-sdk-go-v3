package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateServerRequestServerCertificate 服务端证书。 隧道协议类型是SSL时，必填
type CreateServerRequestServerCertificate struct {

	// 服务端证书ID,为CCM服务中的证书ID。服务端证书类型为CCM时，必填
	Id *string `json:"id,omitempty"`

	// 证书来源
	Source *CreateServerRequestServerCertificateSource `json:"source,omitempty"`
}

func (o CreateServerRequestServerCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerRequestServerCertificate struct{}"
	}

	return strings.Join([]string{"CreateServerRequestServerCertificate", string(data)}, " ")
}

type CreateServerRequestServerCertificateSource struct {
	value string
}

type CreateServerRequestServerCertificateSourceEnum struct {
	CCM          CreateServerRequestServerCertificateSource
	SERVICE_SIGN CreateServerRequestServerCertificateSource
}

func GetCreateServerRequestServerCertificateSourceEnum() CreateServerRequestServerCertificateSourceEnum {
	return CreateServerRequestServerCertificateSourceEnum{
		CCM: CreateServerRequestServerCertificateSource{
			value: "CCM",
		},
		SERVICE_SIGN: CreateServerRequestServerCertificateSource{
			value: "SERVICE_SIGN",
		},
	}
}

func (c CreateServerRequestServerCertificateSource) Value() string {
	return c.value
}

func (c CreateServerRequestServerCertificateSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateServerRequestServerCertificateSource) UnmarshalJSON(b []byte) error {
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
