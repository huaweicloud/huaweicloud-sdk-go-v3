package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowIpInfoResponse struct {
	// IP归属信息列表。

	CdnIps         *[]CdnIps `json:"cdn_ips,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowIpInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowIpInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowIpInfoResponse", string(data)}, " ")
}
