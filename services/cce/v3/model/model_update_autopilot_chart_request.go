package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutopilotChartRequest Request Object
type UpdateAutopilotChartRequest struct {

	// **参数解释：** 模板的ID。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ChartId string `json:"chart_id"`

	Body *UpdateAutopilotChartRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UpdateAutopilotChartRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutopilotChartRequest struct{}"
	}

	return strings.Join([]string{"UpdateAutopilotChartRequest", string(data)}, " ")
}
