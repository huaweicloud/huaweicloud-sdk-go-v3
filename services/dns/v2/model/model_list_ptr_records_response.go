package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPtrRecordsResponse struct {
	Links          *PageLink                     `json:"links,omitempty"`
	Metadata       *Metedata                     `json:"metadata,omitempty"`
	Floatingips    *[]ListPtrRecordsFloatingResp `json:"floatingips,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListPtrRecordsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPtrRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListPtrRecordsResponse", string(data)}, " ")
}
