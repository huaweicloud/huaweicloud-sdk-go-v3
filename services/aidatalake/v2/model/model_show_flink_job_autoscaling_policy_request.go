package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkJobAutoscalingPolicyRequest Request Object
type ShowFlinkJobAutoscalingPolicyRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：Flink作业ID，用于唯一标识一个Flink作业，可通过[查询Flink作业列表](ListFlinkJobs.xml)接口获取。 **约束限制**：不涉及。 **取值范围**：采用UUIDv7格式，固定36个字符。 **默认取值**：不涉及。
	JobId string `json:"job_id"`
}

func (o ShowFlinkJobAutoscalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkJobAutoscalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowFlinkJobAutoscalingPolicyRequest", string(data)}, " ")
}
