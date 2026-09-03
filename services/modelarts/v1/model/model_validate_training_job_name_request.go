package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateTrainingJobNameRequest Request Object
type ValidateTrainingJobNameRequest struct {

	// **参数解释**：训练作业名称。 **约束限制**：1 - 64字符，字母、数字、下划线和中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	JobName string `json:"job_name"`

	// **参数解释**：工作空间ID。获取方法请参见[[查询工作空间列表](ListWorkspace.xml)](tag:hc,hk)。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o ValidateTrainingJobNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateTrainingJobNameRequest struct{}"
	}

	return strings.Join([]string{"ValidateTrainingJobNameRequest", string(data)}, " ")
}
