package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineLogResponse Response Object
type ShowPipelineLogResponse struct {

	// **参数解释**： 是否有更多日志。 **取值范围**： - true：有更多日志。 - false：没有更多日志。
	HasMore *bool `json:"has_more,omitempty"`

	// **参数解释**： 查询日志结束偏移。填入请求体end_offset字段，用于查询下一页日志。 **取值范围**： 不涉及。
	EndOffset *string `json:"end_offset,omitempty"`

	// **参数解释**： 查询日志起始偏移。填入请求体start_offset字段，用于查询下一页日志。 **取值范围**： 不涉及。
	StartOffset *string `json:"start_offset,omitempty"`

	// **参数解释**： 日志内容。 **取值范围**： 不涉及。
	Log *string `json:"log,omitempty"`

	// **参数解释**： 所属步骤ID。 **取值范围**： 不涉及。
	StepRunId      *string `json:"step_run_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPipelineLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineLogResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineLogResponse", string(data)}, " ")
}
