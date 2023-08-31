package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseWidgetInfoLocation 监控视图图表坐标
type BaseWidgetInfoLocation struct {

	// 监控视图的上坐标
	Top *int32 `json:"top,omitempty"`

	// 监控视图的左坐标
	Left *int32 `json:"left,omitempty"`

	// 监控视图图表宽度
	Width *int32 `json:"width,omitempty"`

	// 监控视图图表高度
	Height *int32 `json:"height,omitempty"`
}

func (o BaseWidgetInfoLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseWidgetInfoLocation struct{}"
	}

	return strings.Join([]string{"BaseWidgetInfoLocation", string(data)}, " ")
}
