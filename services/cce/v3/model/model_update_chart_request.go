package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateChartRequest Request Object
type UpdateChartRequest struct {

	// **参数解释：** 模板的ID。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ChartId string `json:"chart_id"`

	Body *UpdateChartRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UpdateChartRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateChartRequest struct{}"
	}

	return strings.Join([]string{"UpdateChartRequest", string(data)}, " ")
}
