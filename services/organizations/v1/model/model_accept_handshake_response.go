package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptHandshakeResponse Response Object
type AcceptHandshakeResponse struct {
	Handshake      *HandshakeDto `json:"handshake,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o AcceptHandshakeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptHandshakeResponse struct{}"
	}

	return strings.Join([]string{"AcceptHandshakeResponse", string(data)}, " ")
}
