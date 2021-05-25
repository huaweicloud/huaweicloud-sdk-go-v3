package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDomainRequest struct {
	// 直播域名

	Domain string `json:"domain"`
}

func (o DeleteDomainRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDomainRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainRequest", string(data)}, " ")
}
