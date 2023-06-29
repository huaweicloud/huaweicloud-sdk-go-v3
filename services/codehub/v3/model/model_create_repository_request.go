package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRepositoryRequest Request Object
type CreateRepositoryRequest struct {
	Body *CreateRepoRequest `json:"body,omitempty"`
}

func (o CreateRepositoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepositoryRequest struct{}"
	}

	return strings.Join([]string{"CreateRepositoryRequest", string(data)}, " ")
}
