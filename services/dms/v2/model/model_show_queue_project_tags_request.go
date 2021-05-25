package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowQueueProjectTagsRequest struct {
}

func (o ShowQueueProjectTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQueueProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowQueueProjectTagsRequest", string(data)}, " ")
}
