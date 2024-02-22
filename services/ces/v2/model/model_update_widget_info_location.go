package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWidgetInfoLocation 监控视图图表坐标
type UpdateWidgetInfoLocation struct {

	// 监控视图的上坐标
	Top int32 `json:"top"`

	// 监控视图的左坐标
	Left int32 `json:"left"`

	// 监控视图图表宽度
	Width int32 `json:"width"`

	// 监控视图图表高度
	Height int32 `json:"height"`
}

func (o UpdateWidgetInfoLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWidgetInfoLocation struct{}"
	}

	return strings.Join([]string{"UpdateWidgetInfoLocation", string(data)}, " ")
}
