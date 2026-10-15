package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelSparkJobRequest Request Object
type CancelSparkJobRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：Spark作业ID，用于唯一标识作业。 **约束限制**：不涉及。 **取值范围**：只能由英文字母、数字、下划线和中划线组成，且长度为1~64个字符。 **默认取值**：不涉及。
	JobId string `json:"job_id"`
}

func (o CancelSparkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSparkJobRequest struct{}"
	}

	return strings.Join([]string{"CancelSparkJobRequest", string(data)}, " ")
}
