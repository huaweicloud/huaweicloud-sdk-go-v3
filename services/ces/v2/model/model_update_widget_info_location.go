package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWidgetInfoLocation **参数解释** 监控视图图表坐标 **约束限制** 不涉及
type UpdateWidgetInfoLocation struct {

	// **参数解释** 监控视图的上坐标 **约束限制** 不涉及 **取值范围** 坐标的取值范围[0,2147483647] **默认取值** 不涉及
	Top int32 `json:"top"`

	// **参数解释** 监控视图的左坐标 **约束限制** 不涉及 **取值范围** 坐标的取值范围[0,9] **默认取值** 不涉及
	Left int32 `json:"left"`

	// **参数解释** 监控视图图表宽度 **约束限制** 不涉及 **取值范围** 宽度的取值范围[3,12] **默认取值** 不涉及
	Width int32 `json:"width"`

	// **参数解释** 监控视图图表高度 **约束限制** 不涉及 **取值范围** 高度的取值范围[3,2147483647] **默认取值** 不涉及
	Height int32 `json:"height"`
}

func (o UpdateWidgetInfoLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWidgetInfoLocation struct{}"
	}

	return strings.Join([]string{"UpdateWidgetInfoLocation", string(data)}, " ")
}
