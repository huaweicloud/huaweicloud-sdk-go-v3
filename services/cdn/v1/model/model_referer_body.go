package model

import (
	"encoding/json"

	"strings"
)

type RefererBody struct {
	Referer *Referer `json:"referer"`
}

func (o RefererBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RefererBody struct{}"
	}

	return strings.Join([]string{"RefererBody", string(data)}, " ")
}
