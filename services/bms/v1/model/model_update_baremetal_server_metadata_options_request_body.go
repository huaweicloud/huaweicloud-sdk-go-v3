package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateBaremetalServerMetadataOptionsRequestBody 裸金属服务器元数据配置请求体。
type UpdateBaremetalServerMetadataOptionsRequestBody struct {

	// 是否禁用IMDS服务。
	HttpEndpoint *UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpoint `json:"http_endpoint,omitempty"`

	// 是否要求携带token。
	HttpTokens *UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokens `json:"http_tokens,omitempty"`
}

func (o UpdateBaremetalServerMetadataOptionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBaremetalServerMetadataOptionsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateBaremetalServerMetadataOptionsRequestBody", string(data)}, " ")
}

type UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpoint struct {
	value string
}

type UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpointEnum struct {
	ENABLED  UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpoint
	DISABLED UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpoint
}

func GetUpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpointEnum() UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpointEnum {
	return UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpointEnum{
		ENABLED: UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpoint{
			value: "enabled",
		},
		DISABLED: UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpoint{
			value: "disabled",
		},
	}
}

func (c UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpoint) Value() string {
	return c.value
}

func (c UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpoint) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateBaremetalServerMetadataOptionsRequestBodyHttpEndpoint) UnmarshalJSON(b []byte) error {
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

type UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokens struct {
	value string
}

type UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokensEnum struct {
	OPTIONAL UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokens
	REQUIRED UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokens
}

func GetUpdateBaremetalServerMetadataOptionsRequestBodyHttpTokensEnum() UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokensEnum {
	return UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokensEnum{
		OPTIONAL: UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokens{
			value: "optional",
		},
		REQUIRED: UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokens{
			value: "required",
		},
	}
}

func (c UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokens) Value() string {
	return c.value
}

func (c UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokens) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateBaremetalServerMetadataOptionsRequestBodyHttpTokens) UnmarshalJSON(b []byte) error {
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
