package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResJobsRequest Request Object
type CreateResJobsRequest struct {

	// 资源id
	ResourceId string `json:"resource_id"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *CreateResJobsReququestBody `json:"body,omitempty"`
}

func (o CreateResJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResJobsRequest struct{}"
	}

	return strings.Join([]string{"CreateResJobsRequest", string(data)}, " ")
}
