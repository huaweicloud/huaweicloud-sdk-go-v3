package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NovaListServerSecurityGroupsRequest struct {
	// 云服务器ID。

	ServerId string `json:"server_id"`
}

func (o NovaListServerSecurityGroupsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NovaListServerSecurityGroupsRequest struct{}"
	}

	return strings.Join([]string{"NovaListServerSecurityGroupsRequest", string(data)}, " ")
}
