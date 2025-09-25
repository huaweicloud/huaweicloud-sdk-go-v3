package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChatAddition struct {

	// **参数解释**： 实验问答使用的空间ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**： 实验问答的作业ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	JobId *string `json:"job_id,omitempty"`

	ToolType *ToolType `json:"tool_type,omitempty"`

	// **参数解释**： 作业类型。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	JobType *string `json:"job_type,omitempty"`

	// **参数解释**： 深度探究生成报告存储路径。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	DataPath *string `json:"data_path,omitempty"`
}

func (o ChatAddition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatAddition struct{}"
	}

	return strings.Join([]string{"ChatAddition", string(data)}, " ")
}
