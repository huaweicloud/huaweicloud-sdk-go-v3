package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListClustersByTagsRequest struct {
	Body *ListResourceReq `json:"body,omitempty"`
}

func (o ListClustersByTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListClustersByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListClustersByTagsRequest", string(data)}, " ")
}
