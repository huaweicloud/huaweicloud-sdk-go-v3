package model

import (
	"encoding/json"

	"strings"
)

type AvParameters struct {
	Video  *VideoParameters `json:"video,omitempty"`
	Audio  *Audio           `json:"audio,omitempty"`
	Common *Common          `json:"common"`
}

func (o AvParameters) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AvParameters struct{}"
	}

	return strings.Join([]string{"AvParameters", string(data)}, " ")
}
