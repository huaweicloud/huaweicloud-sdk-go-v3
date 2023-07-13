package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelinesRequest Request Object
type ListPipelinesRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	Body *ListPipelineQuery `json:"body,omitempty"`
}

func (o ListPipelinesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesRequest struct{}"
	}

	return strings.Join([]string{"ListPipelinesRequest", string(data)}, " ")
}
