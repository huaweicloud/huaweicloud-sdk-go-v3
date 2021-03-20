package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListRecordSetsWithLineResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Recordsets *[]QueryRecordSetWithLineResp `json:"recordsets,omitempty"`

	Metadata       *Metedata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsWithLineResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRecordSetsWithLineResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsWithLineResponse", string(data)}, " ")
}
