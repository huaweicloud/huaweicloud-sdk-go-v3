package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type EnableLogCollectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableLogCollectionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnableLogCollectionResponse struct{}"
	}

	return strings.Join([]string{"EnableLogCollectionResponse", string(data)}, " ")
}
