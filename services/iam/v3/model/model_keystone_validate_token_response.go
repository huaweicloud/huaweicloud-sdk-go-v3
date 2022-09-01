package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type KeystoneValidateTokenResponse struct {
	Token *TokenResult `json:"token,omitempty" xml:"token"`

	XSubjectToken  *string `json:"X-Subject-Token,omitempty" xml:"X-Subject-Token"`
	HttpStatusCode int     `json:"-"`
}

func (o KeystoneValidateTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneValidateTokenResponse struct{}"
	}

	return strings.Join([]string{"KeystoneValidateTokenResponse", string(data)}, " ")
}
