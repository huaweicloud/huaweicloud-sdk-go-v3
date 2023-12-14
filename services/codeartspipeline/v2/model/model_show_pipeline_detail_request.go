package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineDetailRequest Request Object
type ShowPipelineDetailRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`
}

func (o ShowPipelineDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowPipelineDetailRequest", string(data)}, " ")
}
