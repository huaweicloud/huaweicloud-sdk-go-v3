package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTokenWithIdTokenResponse Response Object
type CreateTokenWithIdTokenResponse struct {
	Token *ScopedTokenInfo `json:"token,omitempty"`

	XSubjectToken  *string `json:"X-Subject-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTokenWithIdTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTokenWithIdTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateTokenWithIdTokenResponse", string(data)}, " ")
}
