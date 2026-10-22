package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelRayJobRequest Request Object
type CancelRayJobRequest struct {

	// **参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：作业ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36的字母、数字、中划线的组合。 **默认取值**：不涉及。
	JobId string `json:"job_id"`
}

func (o CancelRayJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRayJobRequest struct{}"
	}

	return strings.Join([]string{"CancelRayJobRequest", string(data)}, " ")
}
