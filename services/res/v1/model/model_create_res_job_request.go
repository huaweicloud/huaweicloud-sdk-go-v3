package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResJobRequest Request Object
type CreateResJobRequest struct {

	// 资源id
	ResourceId string `json:"resource_id"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *CreateResJobRequestBody `json:"body,omitempty"`
}

func (o CreateResJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResJobRequest struct{}"
	}

	return strings.Join([]string{"CreateResJobRequest", string(data)}, " ")
}
