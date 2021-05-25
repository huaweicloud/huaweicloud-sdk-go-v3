package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowQueueProjectTagsResponse struct {
	// 标签列表

	Tags           *[]ShowProjectTagsRespTags `json:"tags,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowQueueProjectTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQueueProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowQueueProjectTagsResponse", string(data)}, " ")
}
