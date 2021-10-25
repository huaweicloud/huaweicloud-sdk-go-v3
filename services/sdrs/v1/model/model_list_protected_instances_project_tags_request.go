package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProtectedInstancesProjectTagsRequest struct {
}

func (o ListProtectedInstancesProjectTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProtectedInstancesProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProtectedInstancesProjectTagsRequest", string(data)}, " ")
}
