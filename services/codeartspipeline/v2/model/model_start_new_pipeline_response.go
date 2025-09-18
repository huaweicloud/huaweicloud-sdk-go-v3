package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartNewPipelineResponse Response Object
type StartNewPipelineResponse struct {

	// **参数解释**： 流水线ID。 **取值范围**： 32位字符，由数字和字母组成。
	PipelineId *string `json:"pipeline_id,omitempty"`

	// **参数解释**： 流水线执行ID。 **取值范围**： 32位字符，由数字和字母组成。
	BuildId        *string `json:"build_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartNewPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartNewPipelineResponse struct{}"
	}

	return strings.Join([]string{"StartNewPipelineResponse", string(data)}, " ")
}
