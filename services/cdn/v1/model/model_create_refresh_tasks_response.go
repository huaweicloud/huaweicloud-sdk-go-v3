package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateRefreshTasksResponse struct {
	RefreshTask    *RefreshPreheatingBody `json:"refreshTask,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateRefreshTasksResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRefreshTasksResponse struct{}"
	}

	return strings.Join([]string{"CreateRefreshTasksResponse", string(data)}, " ")
}
