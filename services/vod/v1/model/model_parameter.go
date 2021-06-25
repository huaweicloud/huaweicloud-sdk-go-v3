package model

import (
	"encoding/json"

	"strings"
)

type Parameter struct {
	// 封装格式，可选值：“MP3”。<br/>

	Format *string `json:"format,omitempty"`
}

func (o Parameter) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Parameter struct{}"
	}

	return strings.Join([]string{"Parameter", string(data)}, " ")
}
