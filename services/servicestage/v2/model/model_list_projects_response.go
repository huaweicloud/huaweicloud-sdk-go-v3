package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProjectsResponse struct {
	// 项目列表。
	Projects       *[]Project `json:"projects,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListProjectsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectsResponse", string(data)}, " ")
}
