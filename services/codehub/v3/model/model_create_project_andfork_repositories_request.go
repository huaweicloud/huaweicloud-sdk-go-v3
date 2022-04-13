package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProjectAndforkRepositoriesRequest struct {
	Body *ForkProjectRepoRequest `json:"body,omitempty"`
}

func (o CreateProjectAndforkRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectAndforkRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectAndforkRepositoriesRequest", string(data)}, " ")
}
