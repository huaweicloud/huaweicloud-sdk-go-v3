package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type GlanceDeleteTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceDeleteTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceDeleteTagResponse struct{}"
	}

	return strings.Join([]string{"GlanceDeleteTagResponse", string(data)}, " ")
}
