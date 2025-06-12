package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateServerMetadataOptionsRequestBody 云服务器元数据配置请求体。
type UpdateServerMetadataOptionsRequestBody struct {

	// 是否禁用IMDS服务。
	HttpEndpoint *UpdateServerMetadataOptionsRequestBodyHttpEndpoint `json:"http_endpoint,omitempty"`

	// 是否要求携带token。
	HttpTokens *UpdateServerMetadataOptionsRequestBodyHttpTokens `json:"http_tokens,omitempty"`
}

func (o UpdateServerMetadataOptionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerMetadataOptionsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateServerMetadataOptionsRequestBody", string(data)}, " ")
}

type UpdateServerMetadataOptionsRequestBodyHttpEndpoint struct {
	value string
}

type UpdateServerMetadataOptionsRequestBodyHttpEndpointEnum struct {
	ENABLED  UpdateServerMetadataOptionsRequestBodyHttpEndpoint
	DISABLED UpdateServerMetadataOptionsRequestBodyHttpEndpoint
}

func GetUpdateServerMetadataOptionsRequestBodyHttpEndpointEnum() UpdateServerMetadataOptionsRequestBodyHttpEndpointEnum {
	return UpdateServerMetadataOptionsRequestBodyHttpEndpointEnum{
		ENABLED: UpdateServerMetadataOptionsRequestBodyHttpEndpoint{
			value: "enabled",
		},
		DISABLED: UpdateServerMetadataOptionsRequestBodyHttpEndpoint{
			value: "disabled",
		},
	}
}

func (c UpdateServerMetadataOptionsRequestBodyHttpEndpoint) Value() string {
	return c.value
}

func (c UpdateServerMetadataOptionsRequestBodyHttpEndpoint) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateServerMetadataOptionsRequestBodyHttpEndpoint) UnmarshalJSON(b []byte) error {
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

type UpdateServerMetadataOptionsRequestBodyHttpTokens struct {
	value string
}

type UpdateServerMetadataOptionsRequestBodyHttpTokensEnum struct {
	OPTIONAL UpdateServerMetadataOptionsRequestBodyHttpTokens
	REQUIRED UpdateServerMetadataOptionsRequestBodyHttpTokens
}

func GetUpdateServerMetadataOptionsRequestBodyHttpTokensEnum() UpdateServerMetadataOptionsRequestBodyHttpTokensEnum {
	return UpdateServerMetadataOptionsRequestBodyHttpTokensEnum{
		OPTIONAL: UpdateServerMetadataOptionsRequestBodyHttpTokens{
			value: "optional",
		},
		REQUIRED: UpdateServerMetadataOptionsRequestBodyHttpTokens{
			value: "required",
		},
	}
}

func (c UpdateServerMetadataOptionsRequestBodyHttpTokens) Value() string {
	return c.value
}

func (c UpdateServerMetadataOptionsRequestBodyHttpTokens) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateServerMetadataOptionsRequestBodyHttpTokens) UnmarshalJSON(b []byte) error {
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
