package model

import (
	"encoding/json"

	"strings"
)

// IPv6信息
type Ipv6Bandwidth struct {
	// IPv6共享带宽的ID。

	Id string `json:"id"`
}

func (o Ipv6Bandwidth) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Ipv6Bandwidth struct{}"
	}

	return strings.Join([]string{"Ipv6Bandwidth", string(data)}, " ")
}
