package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRepoRequest Request Object
type CreateRepoRequest struct {
	Body *CreateRepoRequestBody `json:"body,omitempty"`
}

func (o CreateRepoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoRequest struct{}"
	}

	return strings.Join([]string{"CreateRepoRequest", string(data)}, " ")
}
