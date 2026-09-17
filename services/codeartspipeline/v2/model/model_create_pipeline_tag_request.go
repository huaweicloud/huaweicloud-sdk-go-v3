package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineTagRequest Request Object
type CreatePipelineTagRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	Body *CreateTagReq `json:"body,omitempty"`
}

func (o CreatePipelineTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineTagRequest struct{}"
	}

	return strings.Join([]string{"CreatePipelineTagRequest", string(data)}, " ")
}
