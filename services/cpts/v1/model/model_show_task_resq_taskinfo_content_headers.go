package model

import (
	"encoding/json"

	"strings"
)

type ShowTaskResqTaskinfoContentHeaders struct {
	// key

	Key *string `json:"key,omitempty"`
	// value

	Value *string `json:"value,omitempty"`
}

func (o ShowTaskResqTaskinfoContentHeaders) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTaskResqTaskinfoContentHeaders struct{}"
	}

	return strings.Join([]string{"ShowTaskResqTaskinfoContentHeaders", string(data)}, " ")
}
