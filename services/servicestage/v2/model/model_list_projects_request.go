package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProjectsRequest struct {
	// 授权名称。

	XRepoAuth string `json:"X-Repo-Auth"`
	// 组织ID。

	Namespace string `json:"namespace"`
}

func (o ListProjectsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectsRequest", string(data)}, " ")
}
