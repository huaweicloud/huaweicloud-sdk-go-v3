package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteL7ruleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteL7ruleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteL7ruleResponse struct{}"
	}

	return strings.Join([]string{"DeleteL7ruleResponse", string(data)}, " ")
}
