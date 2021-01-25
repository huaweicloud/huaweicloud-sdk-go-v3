package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPrivateZoneNameServerResponse struct {
	Nameservers    *[]PrivateNameServer `json:"nameservers,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowPrivateZoneNameServerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPrivateZoneNameServerResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateZoneNameServerResponse", string(data)}, " ")
}
