package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDomainItemLocationDetailsResponse struct {
	DomainItemLocationDetails *DomainItemLocationDetails `json:"domain_item_location_details,omitempty"`
	HttpStatusCode            int                        `json:"-"`
}

func (o ShowDomainItemLocationDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainItemLocationDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainItemLocationDetailsResponse", string(data)}, " ")
}
