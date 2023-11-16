package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineGroupTreeRequest Request Object
type ShowPipelineGroupTreeRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`
}

func (o ShowPipelineGroupTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineGroupTreeRequest struct{}"
	}

	return strings.Join([]string{"ShowPipelineGroupTreeRequest", string(data)}, " ")
}
