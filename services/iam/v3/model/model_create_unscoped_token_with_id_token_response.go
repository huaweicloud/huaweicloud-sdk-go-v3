package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateUnscopedTokenWithIdTokenResponse struct {
	Token *UnscopedTokenInfo `json:"token,omitempty" xml:"token"`

	XSubjectToken  *string `json:"X-Subject-Token,omitempty" xml:"X-Subject-Token"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUnscopedTokenWithIdTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUnscopedTokenWithIdTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateUnscopedTokenWithIdTokenResponse", string(data)}, " ")
}
