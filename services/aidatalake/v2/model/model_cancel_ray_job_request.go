package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelRayJobRequest Request Object
type CancelRayJobRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
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
