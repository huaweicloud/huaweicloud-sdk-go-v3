package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCommitsRequest struct {
	XRepoAuth string  `json:"X-Repo-Auth"`
	Namespace string  `json:"namespace"`
	Project   string  `json:"project"`
	Ref       *string `json:"ref,omitempty"`
}

func (o ListCommitsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCommitsRequest struct{}"
	}

	return strings.Join([]string{"ListCommitsRequest", string(data)}, " ")
}
