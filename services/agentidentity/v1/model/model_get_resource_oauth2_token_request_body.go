package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GetResourceOauth2TokenRequestBody struct {

	// Additional custom parameters for the authorization request (does not override standard OAuth2 parameters)
	CustomParameters map[string]string `json:"custom_parameters,omitempty"`

	// Opaque string for CSRF protection (returned in callback URL response as standard state parameter)
	CustomState *string `json:"custom_state,omitempty"`

	// Whether to initiate a new 3LO flow regardless of existing sessions
	ForceAuthentication *bool `json:"force_authentication,omitempty"`

	// Type of OAuth2 flow to perform
	Oauth2Flow GetResourceOauth2TokenRequestBodyOauth2Flow `json:"oauth2_flow"`

	// Name of the resource's credential provider
	ResourceCredentialProviderName string `json:"resource_credential_provider_name"`

	// Callback URL to redirect after token retrieval (must be configured for workload identity)
	ResourceOauth2ReturnUrl *string `json:"resource_oauth2_return_url,omitempty"`

	// OAuth scopes being requested (coarse-grained permissions, supplemented by rich_authorization_details for fine-grained control)
	Scopes *[]string `json:"scopes,omitempty"`

	// Unique identifier for the user's authentication session (tracks OAuth2 flow state)
	SessionUri *string `json:"session_uri,omitempty"`

	// Identity token of the workload requesting the OAuth2 token
	WorkloadAccessToken *string `json:"workload_access_token,omitempty"`
}

func (o GetResourceOauth2TokenRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceOauth2TokenRequestBody struct{}"
	}

	return strings.Join([]string{"GetResourceOauth2TokenRequestBody", string(data)}, " ")
}

type GetResourceOauth2TokenRequestBodyOauth2Flow struct {
	value string
}

type GetResourceOauth2TokenRequestBodyOauth2FlowEnum struct {
	USER_FEDERATION GetResourceOauth2TokenRequestBodyOauth2Flow
	M2_M            GetResourceOauth2TokenRequestBodyOauth2Flow
}

func GetGetResourceOauth2TokenRequestBodyOauth2FlowEnum() GetResourceOauth2TokenRequestBodyOauth2FlowEnum {
	return GetResourceOauth2TokenRequestBodyOauth2FlowEnum{
		USER_FEDERATION: GetResourceOauth2TokenRequestBodyOauth2Flow{
			value: "USER_FEDERATION",
		},
		M2_M: GetResourceOauth2TokenRequestBodyOauth2Flow{
			value: "M2M",
		},
	}
}

func (c GetResourceOauth2TokenRequestBodyOauth2Flow) Value() string {
	return c.value
}

func (c GetResourceOauth2TokenRequestBodyOauth2Flow) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetResourceOauth2TokenRequestBodyOauth2Flow) UnmarshalJSON(b []byte) error {
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
