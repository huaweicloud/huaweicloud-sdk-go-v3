package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPublicipResponse struct {
	Publicip       *PublicipShowResp `json:"publicip,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowPublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicipResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicipResponse", string(data)}, " ")
}
