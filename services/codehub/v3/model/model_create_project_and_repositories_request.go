package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProjectAndRepositoriesRequest struct {
	Body *CreateProjectRepoRequest `json:"body,omitempty"`
}

func (o CreateProjectAndRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectAndRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectAndRepositoriesRequest", string(data)}, " ")
}
