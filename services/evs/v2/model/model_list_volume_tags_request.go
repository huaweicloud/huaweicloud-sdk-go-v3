package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListVolumeTagsRequest struct {
}

func (o ListVolumeTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"ListVolumeTagsRequest", string(data)}, " ")
}
