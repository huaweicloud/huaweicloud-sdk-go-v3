package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseWidgetInfoRespLocation **参数解释** 监控视图图表坐标
type BaseWidgetInfoRespLocation struct {

	// **参数解释** 监控视图的上坐标              **取值范围** 最小值为0，最大值为2147483647
	Top int32 `json:"top"`

	// **参数解释** 监控视图的左坐标              **取值范围** 最小值为0，最大值为9
	Left int32 `json:"left"`

	// **参数解释** 监控视图图表宽度               **取值范围** 最小值为3，最大值为12
	Width int32 `json:"width"`

	// **参数解释** 监控视图图表高度             **取值范围** 最小值为3，最大值为2147483647
	Height int32 `json:"height"`
}

func (o BaseWidgetInfoRespLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseWidgetInfoRespLocation struct{}"
	}

	return strings.Join([]string{"BaseWidgetInfoRespLocation", string(data)}, " ")
}
