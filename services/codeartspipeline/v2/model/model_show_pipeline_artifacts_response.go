package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineArtifactsResponse Response Object
type ShowPipelineArtifactsResponse struct {

	// **参数解释**： 流水线任务产物列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Artifacts      *[]Artifact `json:"artifacts,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowPipelineArtifactsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineArtifactsResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineArtifactsResponse", string(data)}, " ")
}
