package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CredentialProviderVendor Supported OAuth2 provider vendors.
type CredentialProviderVendor struct {
	value string
}

type CredentialProviderVendorEnum struct {
	MICROSOFT_OAUTH2 CredentialProviderVendor
	GOOGLE_OAUTH2    CredentialProviderVendor
	GITHUB_OAUTH2    CredentialProviderVendor
	CUSTOM_OAUTH2    CredentialProviderVendor
}

func GetCredentialProviderVendorEnum() CredentialProviderVendorEnum {
	return CredentialProviderVendorEnum{
		MICROSOFT_OAUTH2: CredentialProviderVendor{
			value: "MicrosoftOauth2",
		},
		GOOGLE_OAUTH2: CredentialProviderVendor{
			value: "GoogleOauth2",
		},
		GITHUB_OAUTH2: CredentialProviderVendor{
			value: "GithubOauth2",
		},
		CUSTOM_OAUTH2: CredentialProviderVendor{
			value: "CustomOauth2",
		},
	}
}

func (c CredentialProviderVendor) Value() string {
	return c.value
}

func (c CredentialProviderVendor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CredentialProviderVendor) UnmarshalJSON(b []byte) error {
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
