package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDockerRepositoriesRequest Request Object
type CreateDockerRepositoriesRequest struct {
	Body *CreateDockerRepositoryDo `json:"body,omitempty"`
}

func (o CreateDockerRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDockerRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"CreateDockerRepositoriesRequest", string(data)}, " ")
}
