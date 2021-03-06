package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowListenerResponse struct {
	Listener       *ListenerResp `json:"listener,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowListenerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowListenerResponse struct{}"
	}

	return strings.Join([]string{"ShowListenerResponse", string(data)}, " ")
}
