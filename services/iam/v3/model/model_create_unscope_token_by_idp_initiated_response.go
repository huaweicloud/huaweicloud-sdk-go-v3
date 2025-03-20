package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUnscopeTokenByIdpInitiatedResponse Response Object
type CreateUnscopeTokenByIdpInitiatedResponse struct {
	Token *IdpToken `json:"token,omitempty"`

	XSubjectToken  *string `json:"X-Subject-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUnscopeTokenByIdpInitiatedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUnscopeTokenByIdpInitiatedResponse struct{}"
	}

	return strings.Join([]string{"CreateUnscopeTokenByIdpInitiatedResponse", string(data)}, " ")
}
