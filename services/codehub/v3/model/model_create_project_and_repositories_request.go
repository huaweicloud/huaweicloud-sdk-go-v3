package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProjectAndRepositoriesRequest struct {
	Body *CreateProjectRepoRequest `json:"body,omitempty" xml:"body"`
}

func (o CreateProjectAndRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectAndRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectAndRepositoriesRequest", string(data)}, " ")
}
