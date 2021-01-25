package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RestorePtrRecordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestorePtrRecordResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestorePtrRecordResponse struct{}"
	}

	return strings.Join([]string{"RestorePtrRecordResponse", string(data)}, " ")
}
