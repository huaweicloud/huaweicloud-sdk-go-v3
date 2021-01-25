package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DisassociateDomainV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateDomainV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateDomainV2Response struct{}"
	}

	return strings.Join([]string{"DisassociateDomainV2Response", string(data)}, " ")
}
