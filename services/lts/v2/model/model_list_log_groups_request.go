package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListLogGroupsRequest struct {
}

func (o ListLogGroupsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListLogGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListLogGroupsRequest", string(data)}, " ")
}
