package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Oauth2AuthorizeRequest Request Object
type Oauth2AuthorizeRequest struct {

	// OAuth 2.0 PAR standard request URI, references authorization parameters for the OAuth2 flow
	RequestUri string `json:"request_uri"`
}

func (o Oauth2AuthorizeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Oauth2AuthorizeRequest struct{}"
	}

	return strings.Join([]string{"Oauth2AuthorizeRequest", string(data)}, " ")
}
