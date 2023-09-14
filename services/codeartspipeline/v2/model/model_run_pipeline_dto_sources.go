package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunPipelineDtoSources struct {

	// 源类型
	Type string `json:"type"`

	Params *RunPipelineDtoParams `json:"params,omitempty"`
}

func (o RunPipelineDtoSources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineDtoSources struct{}"
	}

	return strings.Join([]string{"RunPipelineDtoSources", string(data)}, " ")
}
