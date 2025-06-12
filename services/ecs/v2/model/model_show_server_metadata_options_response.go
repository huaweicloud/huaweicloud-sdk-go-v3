package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowServerMetadataOptionsResponse Response Object
type ShowServerMetadataOptionsResponse struct {

	// 是否禁用IMDS服务。
	HttpEndpoint *ShowServerMetadataOptionsResponseHttpEndpoint `json:"http_endpoint,omitempty"`

	// 是否要求携带token。
	HttpTokens     *ShowServerMetadataOptionsResponseHttpTokens `json:"http_tokens,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o ShowServerMetadataOptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerMetadataOptionsResponse struct{}"
	}

	return strings.Join([]string{"ShowServerMetadataOptionsResponse", string(data)}, " ")
}

type ShowServerMetadataOptionsResponseHttpEndpoint struct {
	value string
}

type ShowServerMetadataOptionsResponseHttpEndpointEnum struct {
	ENABLED  ShowServerMetadataOptionsResponseHttpEndpoint
	DISABLED ShowServerMetadataOptionsResponseHttpEndpoint
}

func GetShowServerMetadataOptionsResponseHttpEndpointEnum() ShowServerMetadataOptionsResponseHttpEndpointEnum {
	return ShowServerMetadataOptionsResponseHttpEndpointEnum{
		ENABLED: ShowServerMetadataOptionsResponseHttpEndpoint{
			value: "enabled",
		},
		DISABLED: ShowServerMetadataOptionsResponseHttpEndpoint{
			value: "disabled",
		},
	}
}

func (c ShowServerMetadataOptionsResponseHttpEndpoint) Value() string {
	return c.value
}

func (c ShowServerMetadataOptionsResponseHttpEndpoint) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowServerMetadataOptionsResponseHttpEndpoint) UnmarshalJSON(b []byte) error {
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

type ShowServerMetadataOptionsResponseHttpTokens struct {
	value string
}

type ShowServerMetadataOptionsResponseHttpTokensEnum struct {
	OPTIONAL ShowServerMetadataOptionsResponseHttpTokens
	REQUIRED ShowServerMetadataOptionsResponseHttpTokens
}

func GetShowServerMetadataOptionsResponseHttpTokensEnum() ShowServerMetadataOptionsResponseHttpTokensEnum {
	return ShowServerMetadataOptionsResponseHttpTokensEnum{
		OPTIONAL: ShowServerMetadataOptionsResponseHttpTokens{
			value: "optional",
		},
		REQUIRED: ShowServerMetadataOptionsResponseHttpTokens{
			value: "required",
		},
	}
}

func (c ShowServerMetadataOptionsResponseHttpTokens) Value() string {
	return c.value
}

func (c ShowServerMetadataOptionsResponseHttpTokens) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowServerMetadataOptionsResponseHttpTokens) UnmarshalJSON(b []byte) error {
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
