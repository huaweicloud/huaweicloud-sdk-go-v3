package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DescribeApplicationProviderResponse Response Object
type DescribeApplicationProviderResponse struct {

	// 应用程序提供者URN
	ApplicationProviderUrn *string `json:"application_provider_urn,omitempty"`

	DisplayData *DisplayDataDto `json:"display_data,omitempty"`

	// 支持的协议
	FederationProtocol *DescribeApplicationProviderResponseFederationProtocol `json:"federation_protocol,omitempty"`

	// 应用程序提供者唯一标识符（ID）
	ApplicationProviderId *string `json:"application_provider_id,omitempty"`
	HttpStatusCode        int     `json:"-"`
}

func (o DescribeApplicationProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeApplicationProviderResponse struct{}"
	}

	return strings.Join([]string{"DescribeApplicationProviderResponse", string(data)}, " ")
}

type DescribeApplicationProviderResponseFederationProtocol struct {
	value string
}

type DescribeApplicationProviderResponseFederationProtocolEnum struct {
	SAML  DescribeApplicationProviderResponseFederationProtocol
	OAUTH DescribeApplicationProviderResponseFederationProtocol
}

func GetDescribeApplicationProviderResponseFederationProtocolEnum() DescribeApplicationProviderResponseFederationProtocolEnum {
	return DescribeApplicationProviderResponseFederationProtocolEnum{
		SAML: DescribeApplicationProviderResponseFederationProtocol{
			value: "SAML",
		},
		OAUTH: DescribeApplicationProviderResponseFederationProtocol{
			value: "OAUTH",
		},
	}
}

func (c DescribeApplicationProviderResponseFederationProtocol) Value() string {
	return c.value
}

func (c DescribeApplicationProviderResponseFederationProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DescribeApplicationProviderResponseFederationProtocol) UnmarshalJSON(b []byte) error {
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
