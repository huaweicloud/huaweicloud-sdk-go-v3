package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPipelineResponse Response Object
type RunPipelineResponse struct {

	// **参数解释**： 流水线运行实例ID，[启动流水线](RunPipeline.xml)接口的返回值即为流水线运行实例ID。 **取值范围**： 32位字符，仅由数字和字母组成。
	PipelineRunId  *string `json:"pipeline_run_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineResponse struct{}"
	}

	return strings.Join([]string{"RunPipelineResponse", string(data)}, " ")
}
