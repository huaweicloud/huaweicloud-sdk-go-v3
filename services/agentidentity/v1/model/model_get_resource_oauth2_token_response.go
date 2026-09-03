package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// GetResourceOauth2TokenResponse Response Object
type GetResourceOauth2TokenResponse struct {

	// OAuth2.0 access token to use (embedded with RAR authorization details claims if rich_authorization_details were requested)
	AccessToken *string `json:"access_token,omitempty"`

	// URL to initiate authorization (provided when user authorization is required, includes encoded RAR details for user consent)
	AuthorizationUrl *string `json:"authorization_url,omitempty"`

	// Status of the user's authorization session (determines next steps in OAuth2 flow)
	SessionStatus *GetResourceOauth2TokenResponseSessionStatus `json:"session_status,omitempty"`

	// Unique identifier for the user's authentication session (matches request session_uri)
	SessionUri *string `json:"session_uri,omitempty"`

	// Absolute expiration time of the access token in RFC 3339 format, UTC timezone.
	ExpiresAt      *sdktime.SdkTime `json:"expires_at,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o GetResourceOauth2TokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceOauth2TokenResponse struct{}"
	}

	return strings.Join([]string{"GetResourceOauth2TokenResponse", string(data)}, " ")
}

type GetResourceOauth2TokenResponseSessionStatus struct {
	value string
}

type GetResourceOauth2TokenResponseSessionStatusEnum struct {
	IN_PROGRESS GetResourceOauth2TokenResponseSessionStatus
	FAILED      GetResourceOauth2TokenResponseSessionStatus
}

func GetGetResourceOauth2TokenResponseSessionStatusEnum() GetResourceOauth2TokenResponseSessionStatusEnum {
	return GetResourceOauth2TokenResponseSessionStatusEnum{
		IN_PROGRESS: GetResourceOauth2TokenResponseSessionStatus{
			value: "IN_PROGRESS",
		},
		FAILED: GetResourceOauth2TokenResponseSessionStatus{
			value: "FAILED",
		},
	}
}

func (c GetResourceOauth2TokenResponseSessionStatus) Value() string {
	return c.value
}

func (c GetResourceOauth2TokenResponseSessionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetResourceOauth2TokenResponseSessionStatus) UnmarshalJSON(b []byte) error {
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
