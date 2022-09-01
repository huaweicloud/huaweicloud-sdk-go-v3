package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateWebSocketTokenResponse struct {

	// websocket建链Token。
	WebSocketToken *string `json:"webSocketToken,omitempty" xml:"webSocketToken"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWebSocketTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWebSocketTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateWebSocketTokenResponse", string(data)}, " ")
}
