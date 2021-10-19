package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AcceptOrRejectEndpointResponse struct {
	// 连接列表

	Connections    *[]Endpoints `json:"connections,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o AcceptOrRejectEndpointResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AcceptOrRejectEndpointResponse struct{}"
	}

	return strings.Join([]string{"AcceptOrRejectEndpointResponse", string(data)}, " ")
}
