package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateApiTestSuiteByRepoFileRequest struct {
	Body *CreateTestSuitByRepoFileInfo `json:"body,omitempty"`
}

func (o CreateApiTestSuiteByRepoFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiTestSuiteByRepoFileRequest struct{}"
	}

	return strings.Join([]string{"CreateApiTestSuiteByRepoFileRequest", string(data)}, " ")
}
