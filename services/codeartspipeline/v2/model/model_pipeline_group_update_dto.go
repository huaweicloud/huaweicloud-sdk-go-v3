package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineGroupUpdateDto PipelineGroupUpdateDTO
type PipelineGroupUpdateDto struct {

	// 流水线分组名
	Name string `json:"name"`

	// 流水线分组ID
	Id string `json:"id"`
}

func (o PipelineGroupUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineGroupUpdateDto struct{}"
	}

	return strings.Join([]string{"PipelineGroupUpdateDto", string(data)}, " ")
}
