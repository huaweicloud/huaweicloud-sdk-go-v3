package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListFlavorsRequest struct {
}

func (o ListFlavorsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}
