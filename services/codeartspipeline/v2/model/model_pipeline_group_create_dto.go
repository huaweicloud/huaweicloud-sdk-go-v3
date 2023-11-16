package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineGroupCreateDto PipelineGroupCreateDTO
type PipelineGroupCreateDto struct {

	// 流水线分组名
	Name string `json:"name"`

	// 项目名
	ProjectId string `json:"project_id"`

	// 父分组ID
	ParentId *string `json:"parent_id,omitempty"`
}

func (o PipelineGroupCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineGroupCreateDto struct{}"
	}

	return strings.Join([]string{"PipelineGroupCreateDto", string(data)}, " ")
}
