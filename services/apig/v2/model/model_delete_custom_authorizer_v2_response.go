package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteCustomAuthorizerV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomAuthorizerV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomAuthorizerV2Response struct{}"
	}

	return strings.Join([]string{"DeleteCustomAuthorizerV2Response", string(data)}, " ")
}
