package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPrivateZonesResponse struct {
	Links          *PageLink          `json:"links,omitempty"`
	Metadata       *Metedata          `json:"metadata,omitempty"`
	Zones          *[]PrivateZoneResp `json:"zones,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListPrivateZonesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPrivateZonesResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateZonesResponse", string(data)}, " ")
}
