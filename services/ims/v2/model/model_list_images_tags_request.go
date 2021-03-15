package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListImagesTagsRequest struct {
}

func (o ListImagesTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListImagesTagsRequest struct{}"
	}

	return strings.Join([]string{"ListImagesTagsRequest", string(data)}, " ")
}
