package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineLatestRunArtifactParams **参数解释**： 制品源参数。 **取值范围**： 不涉及。
type PipelineLatestRunArtifactParams struct {

	// **参数解释**： 包版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 过滤分支。 **取值范围**： 不涉及。
	BranchFilter *string `json:"branch_filter,omitempty"`

	// **参数解释**： 包名称。 **取值范围**： 不涉及。
	PackageName *string `json:"package_name,omitempty"`

	// **参数解释**： docker组织。 **取值范围**： 不涉及。
	Organization *string `json:"organization,omitempty"`
}

func (o PipelineLatestRunArtifactParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineLatestRunArtifactParams struct{}"
	}

	return strings.Join([]string{"PipelineLatestRunArtifactParams", string(data)}, " ")
}
