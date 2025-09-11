package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseWidgetInfoLocation **参数解释** 监控视图图表坐标 **约束限制** 不涉及
type BaseWidgetInfoLocation struct {

	// **参数解释** 监控视图的上坐标 **约束限制** 不涉及               **取值范围** 最小值为0，最大值为2147483647 **默认取值** 不涉及
	Top int32 `json:"top"`

	// **参数解释** 监控视图的左坐标 **约束限制** 不涉及               **取值范围** 最小值为0，最大值为9 **默认取值** 不涉及
	Left int32 `json:"left"`

	// **参数解释** 监控视图图表宽度 **约束限制** 不涉及               **取值范围** 最小值为3，最大值为12 **默认取值** 不涉及
	Width int32 `json:"width"`

	// **参数解释** 监控视图图表高度 **约束限制** 不涉及               **取值范围** 最小值为3，最大值为2147483647 **默认取值** 不涉及
	Height int32 `json:"height"`
}

func (o BaseWidgetInfoLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseWidgetInfoLocation struct{}"
	}

	return strings.Join([]string{"BaseWidgetInfoLocation", string(data)}, " ")
}
