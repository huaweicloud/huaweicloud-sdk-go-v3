package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowListenerResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	Listener       *Listener `json:"listener,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowListenerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowListenerResponse struct{}"
	}

	return strings.Join([]string{"ShowListenerResponse", string(data)}, " ")
}
