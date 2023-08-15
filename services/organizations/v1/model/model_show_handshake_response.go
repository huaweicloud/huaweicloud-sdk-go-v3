package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHandshakeResponse Response Object
type ShowHandshakeResponse struct {
	Handshake      *HandshakeDto `json:"handshake,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowHandshakeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHandshakeResponse struct{}"
	}

	return strings.Join([]string{"ShowHandshakeResponse", string(data)}, " ")
}
