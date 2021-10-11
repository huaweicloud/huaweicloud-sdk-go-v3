package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RemoveEipV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveEipV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveEipV2Response struct{}"
	}

	return strings.Join([]string{"RemoveEipV2Response", string(data)}, " ")
}
