package model

import (
	"encoding/json"

	"strings"
)

//
type OsfederationProtocol struct {
	// 协议ID。

	Id string `json:"id"`
}

func (o OsfederationProtocol) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OsfederationProtocol struct{}"
	}

	return strings.Join([]string{"OsfederationProtocol", string(data)}, " ")
}
