package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAutopilotChartRequest Request Object
type DeleteAutopilotChartRequest struct {

	// **参数解释：** 模板的ID。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ChartId string `json:"chart_id"`
}

func (o DeleteAutopilotChartRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAutopilotChartRequest struct{}"
	}

	return strings.Join([]string{"DeleteAutopilotChartRequest", string(data)}, " ")
}
