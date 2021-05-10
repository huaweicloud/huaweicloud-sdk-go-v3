package model

import (
	"encoding/json"

	"strings"
)

type CreateTagReq struct {
	Tag *Tag `json:"tag"`
}

func (o CreateTagReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTagReq struct{}"
	}

	return strings.Join([]string{"CreateTagReq", string(data)}, " ")
}
