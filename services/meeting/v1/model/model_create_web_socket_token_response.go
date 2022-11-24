package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateWebSocketTokenResponse struct {

	// WebSocket建链Token(有效期1分钟，且一次有效)。
	WebSocketToken *string `json:"webSocketToken,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWebSocketTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWebSocketTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateWebSocketTokenResponse", string(data)}, " ")
}
