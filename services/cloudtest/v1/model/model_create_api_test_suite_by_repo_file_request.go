package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateApiTestSuiteByRepoFileRequest struct {
	Body *CreateTestSuitByRepoFileInfo `json:"body,omitempty"`
}

func (o CreateApiTestSuiteByRepoFileRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateApiTestSuiteByRepoFileRequest struct{}"
	}

	return strings.Join([]string{"CreateApiTestSuiteByRepoFileRequest", string(data)}, " ")
}
