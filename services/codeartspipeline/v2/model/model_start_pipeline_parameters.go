package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartPipelineParameters struct {

	// **参数解释**： 启动流水线时的构建参数。 **约束限制**： 不涉及。 **取值范围**： 不超过8192个字符。 **默认取值**： 不涉及。
	BuildParams *[]StartPipelineBuildParams `json:"build_params,omitempty"`
}

func (o StartPipelineParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPipelineParameters struct{}"
	}

	return strings.Join([]string{"StartPipelineParameters", string(data)}, " ")
}
