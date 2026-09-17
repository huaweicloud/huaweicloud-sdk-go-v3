package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineTagRequest Request Object
type ListPipelineTagRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 项目ID
	ProjId *string `json:"proj_id,omitempty"`
}

func (o ListPipelineTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineTagRequest struct{}"
	}

	return strings.Join([]string{"ListPipelineTagRequest", string(data)}, " ")
}
