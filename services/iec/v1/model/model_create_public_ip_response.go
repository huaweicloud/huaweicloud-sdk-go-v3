package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePublicIpResponse struct {
	Publicip       *PublicIp `json:"publicip,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreatePublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicIpResponse struct{}"
	}

	return strings.Join([]string{"CreatePublicIpResponse", string(data)}, " ")
}
