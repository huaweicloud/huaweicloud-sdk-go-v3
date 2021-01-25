package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteAnimatedGraphicsTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAnimatedGraphicsTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAnimatedGraphicsTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteAnimatedGraphicsTaskResponse", string(data)}, " ")
}
