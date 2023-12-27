package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectRelatedRepositoryRequest Request Object
type CreateProjectRelatedRepositoryRequest struct {
	Body *IdePrivilageProjectInfo `json:"body,omitempty"`
}

func (o CreateProjectRelatedRepositoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectRelatedRepositoryRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectRelatedRepositoryRequest", string(data)}, " ")
}
