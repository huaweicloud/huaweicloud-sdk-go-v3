package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InviteAccountResponse Response Object
type InviteAccountResponse struct {
	Handshake      *HandshakeDto `json:"handshake,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o InviteAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteAccountResponse struct{}"
	}

	return strings.Join([]string{"InviteAccountResponse", string(data)}, " ")
}
