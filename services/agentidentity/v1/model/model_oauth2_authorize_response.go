package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Oauth2AuthorizeResponse Response Object
type Oauth2AuthorizeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o Oauth2AuthorizeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Oauth2AuthorizeResponse struct{}"
	}

	return strings.Join([]string{"Oauth2AuthorizeResponse", string(data)}, " ")
}
