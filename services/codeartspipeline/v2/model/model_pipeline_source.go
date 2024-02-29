package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineSource struct {

	// 流水线源类型
	Type *string `json:"type,omitempty"`

	Params *PipelineSourceParam `json:"params,omitempty"`
}

func (o PipelineSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineSource struct{}"
	}

	return strings.Join([]string{"PipelineSource", string(data)}, " ")
}
