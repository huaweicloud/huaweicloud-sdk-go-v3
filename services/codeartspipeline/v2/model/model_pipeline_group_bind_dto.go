package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineGroupBindDto PipelineGroupBindDTO
type PipelineGroupBindDto struct {

	// 分组ID
	GroupId string `json:"group_id"`

	// 流水线集合
	Pipelines []PipelineGroupBindDtoPipelines `json:"pipelines"`
}

func (o PipelineGroupBindDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineGroupBindDto struct{}"
	}

	return strings.Join([]string{"PipelineGroupBindDto", string(data)}, " ")
}
