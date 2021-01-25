package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type GlanceDeleteImageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceDeleteImageResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageResponse struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageResponse", string(data)}, " ")
}
