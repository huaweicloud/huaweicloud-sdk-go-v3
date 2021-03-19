package model

import (
	"encoding/json"

	"strings"
)

type RestoreInstanceRequestBody struct {
	Source *Source `json:"source"`

	Target *Target `json:"target"`
}

func (o RestoreInstanceRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreInstanceRequestBody", string(data)}, " ")
}
