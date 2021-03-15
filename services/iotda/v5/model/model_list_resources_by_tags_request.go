package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListResourcesByTagsRequest struct {
	InstanceId *string                 `json:"Instance-Id,omitempty"`
	Limit      *int32                  `json:"limit,omitempty"`
	Marker     *string                 `json:"marker,omitempty"`
	Offset     *int32                  `json:"offset,omitempty"`
	Body       *QueryResourceByTagsDto `json:"body,omitempty"`
}

func (o ListResourcesByTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsRequest", string(data)}, " ")
}
