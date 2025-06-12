package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowBaremetalServerMetadataOptionsResponse Response Object
type ShowBaremetalServerMetadataOptionsResponse struct {

	// 是否禁用IMDS服务。
	HttpEndpoint *ShowBaremetalServerMetadataOptionsResponseHttpEndpoint `json:"http_endpoint,omitempty"`

	// 是否要求携带token。
	HttpTokens     *ShowBaremetalServerMetadataOptionsResponseHttpTokens `json:"http_tokens,omitempty"`
	HttpStatusCode int                                                   `json:"-"`
}

func (o ShowBaremetalServerMetadataOptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaremetalServerMetadataOptionsResponse struct{}"
	}

	return strings.Join([]string{"ShowBaremetalServerMetadataOptionsResponse", string(data)}, " ")
}

type ShowBaremetalServerMetadataOptionsResponseHttpEndpoint struct {
	value string
}

type ShowBaremetalServerMetadataOptionsResponseHttpEndpointEnum struct {
	ENABLED  ShowBaremetalServerMetadataOptionsResponseHttpEndpoint
	DISABLED ShowBaremetalServerMetadataOptionsResponseHttpEndpoint
}

func GetShowBaremetalServerMetadataOptionsResponseHttpEndpointEnum() ShowBaremetalServerMetadataOptionsResponseHttpEndpointEnum {
	return ShowBaremetalServerMetadataOptionsResponseHttpEndpointEnum{
		ENABLED: ShowBaremetalServerMetadataOptionsResponseHttpEndpoint{
			value: "enabled",
		},
		DISABLED: ShowBaremetalServerMetadataOptionsResponseHttpEndpoint{
			value: "disabled",
		},
	}
}

func (c ShowBaremetalServerMetadataOptionsResponseHttpEndpoint) Value() string {
	return c.value
}

func (c ShowBaremetalServerMetadataOptionsResponseHttpEndpoint) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBaremetalServerMetadataOptionsResponseHttpEndpoint) UnmarshalJSON(b []byte) error {
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

type ShowBaremetalServerMetadataOptionsResponseHttpTokens struct {
	value string
}

type ShowBaremetalServerMetadataOptionsResponseHttpTokensEnum struct {
	OPTIONAL ShowBaremetalServerMetadataOptionsResponseHttpTokens
	REQUIRED ShowBaremetalServerMetadataOptionsResponseHttpTokens
}

func GetShowBaremetalServerMetadataOptionsResponseHttpTokensEnum() ShowBaremetalServerMetadataOptionsResponseHttpTokensEnum {
	return ShowBaremetalServerMetadataOptionsResponseHttpTokensEnum{
		OPTIONAL: ShowBaremetalServerMetadataOptionsResponseHttpTokens{
			value: "optional",
		},
		REQUIRED: ShowBaremetalServerMetadataOptionsResponseHttpTokens{
			value: "required",
		},
	}
}

func (c ShowBaremetalServerMetadataOptionsResponseHttpTokens) Value() string {
	return c.value
}

func (c ShowBaremetalServerMetadataOptionsResponseHttpTokens) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBaremetalServerMetadataOptionsResponseHttpTokens) UnmarshalJSON(b []byte) error {
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
