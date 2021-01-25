package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListMembersRequest struct {
}

func (o ListMembersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMembersRequest struct{}"
	}

	return strings.Join([]string{"ListMembersRequest", string(data)}, " ")
}
