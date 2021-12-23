package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePublicIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicIpResponse struct{}"
	}

	return strings.Join([]string{"DeletePublicIpResponse", string(data)}, " ")
}
