package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowCommandRequest struct {
	// 命令对应的服务器id

	ServerId string `json:"server_id"`
}

func (o ShowCommandRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCommandRequest struct{}"
	}

	return strings.Join([]string{"ShowCommandRequest", string(data)}, " ")
}
