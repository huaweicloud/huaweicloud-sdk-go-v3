package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CancelingAuthorizationV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelingAuthorizationV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelingAuthorizationV2Response struct{}"
	}

	return strings.Join([]string{"CancelingAuthorizationV2Response", string(data)}, " ")
}
