package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelHandshakeResponse Response Object
type CancelHandshakeResponse struct {
	Handshake      *HandshakeDto `json:"handshake,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CancelHandshakeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelHandshakeResponse struct{}"
	}

	return strings.Join([]string{"CancelHandshakeResponse", string(data)}, " ")
}
