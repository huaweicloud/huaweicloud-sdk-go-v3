package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMavenRepoRequest Request Object
type CreateMavenRepoRequest struct {
	Body *IdeRepositoryDo `json:"body,omitempty"`
}

func (o CreateMavenRepoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMavenRepoRequest struct{}"
	}

	return strings.Join([]string{"CreateMavenRepoRequest", string(data)}, " ")
}
