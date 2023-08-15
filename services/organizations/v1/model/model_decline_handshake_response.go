package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeclineHandshakeResponse Response Object
type DeclineHandshakeResponse struct {
	Handshake      *HandshakeDto `json:"handshake,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o DeclineHandshakeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeclineHandshakeResponse struct{}"
	}

	return strings.Join([]string{"DeclineHandshakeResponse", string(data)}, " ")
}
