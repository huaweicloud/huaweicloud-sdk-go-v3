package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListRecordSetsResponse struct {
	Links          *PageLink         `json:"links,omitempty"`
	Recordsets     *[]ListRecordSets `json:"recordsets,omitempty"`
	Metadata       *Metedata         `json:"metadata,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListRecordSetsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRecordSetsResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsResponse", string(data)}, " ")
}
