package model

import (
	"encoding/json"

	"strings"
)

type VusersBrokens struct {
	// vusers

	Vusers *[]int32 `json:"vusers,omitempty"`
}

func (o VusersBrokens) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VusersBrokens struct{}"
	}

	return strings.Join([]string{"VusersBrokens", string(data)}, " ")
}
