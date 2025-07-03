package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMetadataOptionsResponse Response Object
type ShowMetadataOptionsResponse struct {

	// 是否禁用IMDS服务。
	HttpEndpoint *ShowMetadataOptionsResponseHttpEndpoint `json:"http_endpoint,omitempty"`

	// 是否要求携带token。
	HttpTokens     *ShowMetadataOptionsResponseHttpTokens `json:"http_tokens,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ShowMetadataOptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetadataOptionsResponse struct{}"
	}

	return strings.Join([]string{"ShowMetadataOptionsResponse", string(data)}, " ")
}

type ShowMetadataOptionsResponseHttpEndpoint struct {
	value string
}

type ShowMetadataOptionsResponseHttpEndpointEnum struct {
	ENABLED  ShowMetadataOptionsResponseHttpEndpoint
	DISABLED ShowMetadataOptionsResponseHttpEndpoint
}

func GetShowMetadataOptionsResponseHttpEndpointEnum() ShowMetadataOptionsResponseHttpEndpointEnum {
	return ShowMetadataOptionsResponseHttpEndpointEnum{
		ENABLED: ShowMetadataOptionsResponseHttpEndpoint{
			value: "enabled",
		},
		DISABLED: ShowMetadataOptionsResponseHttpEndpoint{
			value: "disabled",
		},
	}
}

func (c ShowMetadataOptionsResponseHttpEndpoint) Value() string {
	return c.value
}

func (c ShowMetadataOptionsResponseHttpEndpoint) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetadataOptionsResponseHttpEndpoint) UnmarshalJSON(b []byte) error {
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

type ShowMetadataOptionsResponseHttpTokens struct {
	value string
}

type ShowMetadataOptionsResponseHttpTokensEnum struct {
	OPTIONAL ShowMetadataOptionsResponseHttpTokens
	REQUIRED ShowMetadataOptionsResponseHttpTokens
}

func GetShowMetadataOptionsResponseHttpTokensEnum() ShowMetadataOptionsResponseHttpTokensEnum {
	return ShowMetadataOptionsResponseHttpTokensEnum{
		OPTIONAL: ShowMetadataOptionsResponseHttpTokens{
			value: "optional",
		},
		REQUIRED: ShowMetadataOptionsResponseHttpTokens{
			value: "required",
		},
	}
}

func (c ShowMetadataOptionsResponseHttpTokens) Value() string {
	return c.value
}

func (c ShowMetadataOptionsResponseHttpTokens) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetadataOptionsResponseHttpTokens) UnmarshalJSON(b []byte) error {
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
