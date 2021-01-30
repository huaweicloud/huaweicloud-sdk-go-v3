package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RestoreInstanceFromCollectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreInstanceFromCollectionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreInstanceFromCollectionResponse struct{}"
	}

	return strings.Join([]string{"RestoreInstanceFromCollectionResponse", string(data)}, " ")
}
