package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListUserLoginProtectsRequest struct {
}

func (o ListUserLoginProtectsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListUserLoginProtectsRequest struct{}"
	}

	return strings.Join([]string{"ListUserLoginProtectsRequest", string(data)}, " ")
}
