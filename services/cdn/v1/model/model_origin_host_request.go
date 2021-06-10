package model

import (
	"encoding/json"

	"strings"
)

type OriginHostRequest struct {
	OriginHost *OriginHostBody `json:"origin_host"`
}

func (o OriginHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OriginHostRequest struct{}"
	}

	return strings.Join([]string{"OriginHostRequest", string(data)}, " ")
}
