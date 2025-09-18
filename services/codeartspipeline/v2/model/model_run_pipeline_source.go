package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPipelineSource 流水线运行源参数
type RunPipelineSource struct {

	// **参数解释**： 源类型。 **取值范围**： 目前支持code - 代码源类型。
	Type *string `json:"type,omitempty"`

	Params *RunPipelineSourceParams `json:"params,omitempty"`
}

func (o RunPipelineSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineSource struct{}"
	}

	return strings.Join([]string{"RunPipelineSource", string(data)}, " ")
}
