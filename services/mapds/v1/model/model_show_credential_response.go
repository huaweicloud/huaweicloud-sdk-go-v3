package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCredentialResponse Response Object
type ShowCredentialResponse struct {

	// 凭证信息。
	Credentials    *[]CredentialsRespCredentials `json:"credentials,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCredentialResponse struct{}"
	}

	return strings.Join([]string{"ShowCredentialResponse", string(data)}, " ")
}
