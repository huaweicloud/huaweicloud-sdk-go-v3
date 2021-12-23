package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPublicIpResponse struct {
	Publicip       *PublicIp `json:"publicip,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowPublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicIpResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicIpResponse", string(data)}, " ")
}
