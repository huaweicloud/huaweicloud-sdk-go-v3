package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateResJobsRequest struct {

	// 资源id
	ResourceId string `json:"resource_id" xml:"resource_id"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	Body *CreateResJobsReququestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateResJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResJobsRequest struct{}"
	}

	return strings.Join([]string{"CreateResJobsRequest", string(data)}, " ")
}
