package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateProjectRequest struct {
	// 授权名称。

	XRepoAuth string `json:"X-Repo-Auth"`
	// 组织ID。

	Namespace string `json:"namespace"`

	Body *ProjectCreate `json:"body,omitempty"`
}

func (o CreateProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateProjectRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectRequest", string(data)}, " ")
}
