package model

import (
	"encoding/json"

	"strings"
)

type RefreshTaskRequest struct {
	RefreshTask *RefreshTaskRequestRefreshTask `json:"refresh_task"`
}

func (o RefreshTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RefreshTaskRequest struct{}"
	}

	return strings.Join([]string{"RefreshTaskRequest", string(data)}, " ")
}
