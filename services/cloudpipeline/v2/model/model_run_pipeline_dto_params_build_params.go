package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPipelineDtoParamsBuildParams 具体构建参数
type RunPipelineDtoParamsBuildParams struct {

	// 分支还是tag触发
	BuildType *string `json:"build_type,omitempty"`

	// 运行分支
	TargetBranch *string `json:"target_branch,omitempty"`

	// 运行tag
	Tag *string `json:"tag,omitempty"`

	// 触发事件类型
	EventType *string `json:"event_type,omitempty"`
}

func (o RunPipelineDtoParamsBuildParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineDtoParamsBuildParams struct{}"
	}

	return strings.Join([]string{"RunPipelineDtoParamsBuildParams", string(data)}, " ")
}
