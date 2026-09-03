package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingJobLogsFromAomResponse Response Object
type ShowTrainingJobLogsFromAomResponse struct {

	// **参数解释**：返回日志的起始行号。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	StartLine *string `json:"start_line,omitempty"`

	// **参数解释**：返回日志的结束行号。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	EndLine *string `json:"end_line,omitempty"`

	// **参数解释**：返回的日志行数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Lines *int32 `json:"lines,omitempty"`

	// **参数解释**：日志内容。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Content        *string `json:"content,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTrainingJobLogsFromAomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingJobLogsFromAomResponse struct{}"
	}

	return strings.Join([]string{"ShowTrainingJobLogsFromAomResponse", string(data)}, " ")
}
