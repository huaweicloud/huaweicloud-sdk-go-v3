package model

import (
	"encoding/json"

	"strings"
)

type Condition struct {
	PreNodeName *string `json:"preNodeName,omitempty"`
	Expression  *string `json:"expression,omitempty"`
}

func (o Condition) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Condition struct{}"
	}

	return strings.Join([]string{"Condition", string(data)}, " ")
}
