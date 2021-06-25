package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DisableLogCollectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableLogCollectionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisableLogCollectionResponse struct{}"
	}

	return strings.Join([]string{"DisableLogCollectionResponse", string(data)}, " ")
}
