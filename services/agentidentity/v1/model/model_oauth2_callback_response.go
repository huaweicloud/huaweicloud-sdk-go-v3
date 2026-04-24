package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Oauth2CallbackResponse Response Object
type Oauth2CallbackResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o Oauth2CallbackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Oauth2CallbackResponse struct{}"
	}

	return strings.Join([]string{"Oauth2CallbackResponse", string(data)}, " ")
}
