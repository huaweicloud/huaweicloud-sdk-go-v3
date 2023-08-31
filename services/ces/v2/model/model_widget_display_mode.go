package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WidgetDisplayMode struct {

	// 监控视图展示模式，0表示自定义坐标，1代表每行一个
	RowWidgetNum *int32 `json:"row_widget_num,omitempty"`
}

func (o WidgetDisplayMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WidgetDisplayMode struct{}"
	}

	return strings.Join([]string{"WidgetDisplayMode", string(data)}, " ")
}
