package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type GlanceDeleteImageMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceDeleteImageMemberResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageMemberResponse struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageMemberResponse", string(data)}, " ")
}
