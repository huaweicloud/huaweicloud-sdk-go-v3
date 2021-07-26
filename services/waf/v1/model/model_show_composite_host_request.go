package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowCompositeHostRequest struct {
	// 域名Id

	HostId string `json:"host_id"`
}

func (o ShowCompositeHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCompositeHostRequest struct{}"
	}

	return strings.Join([]string{"ShowCompositeHostRequest", string(data)}, " ")
}
