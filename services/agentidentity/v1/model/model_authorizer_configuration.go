package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizerConfiguration Represents inbound authorization configuration options used to authenticate incoming requests.
type AuthorizerConfiguration struct {
	CustomJwt *CustomJwtAuthorizerConfiguration `json:"custom_jwt,omitempty"`

	KeyAuth *KeyAuthAuthorizerConfiguration `json:"key_auth,omitempty"`
}

func (o AuthorizerConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizerConfiguration struct{}"
	}

	return strings.Join([]string{"AuthorizerConfiguration", string(data)}, " ")
}
