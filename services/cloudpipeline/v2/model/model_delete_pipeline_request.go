package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePipelineRequest Request Object
type DeletePipelineRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`
}

func (o DeletePipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineRequest struct{}"
	}

	return strings.Join([]string{"DeletePipelineRequest", string(data)}, " ")
}
