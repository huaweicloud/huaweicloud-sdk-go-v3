package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListBlockchainChannelsResponse struct {
	// 通道信息列表

	Channels       *[]Channel `json:"channels,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListBlockchainChannelsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBlockchainChannelsResponse struct{}"
	}

	return strings.Join([]string{"ListBlockchainChannelsResponse", string(data)}, " ")
}
