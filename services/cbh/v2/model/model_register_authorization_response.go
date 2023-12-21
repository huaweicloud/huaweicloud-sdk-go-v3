package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterAuthorizationResponse Response Object
type RegisterAuthorizationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RegisterAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"RegisterAuthorizationResponse", string(data)}, " ")
}
