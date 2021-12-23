package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartPipelineParameters struct {
	// 启动流水线时的构建参数

	BuildParams *[]StartPipelineBuildParams `json:"build_params,omitempty"`
}

func (o StartPipelineParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPipelineParameters struct{}"
	}

	return strings.Join([]string{"StartPipelineParameters", string(data)}, " ")
}
