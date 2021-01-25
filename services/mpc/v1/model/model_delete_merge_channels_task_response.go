package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteMergeChannelsTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMergeChannelsTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMergeChannelsTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteMergeChannelsTaskResponse", string(data)}, " ")
}
