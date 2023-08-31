package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDashboardRequestBody struct {

	// 自定义监控面板名称
	DashboardName *string `json:"dashboard_name,omitempty"`

	// 监控面板是否标记收藏, true: 收藏, false: 未收藏
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// 监控视图展示模式，0表示自定义坐标，1代表每行一个
	RowWidgetNum *int32 `json:"row_widget_num,omitempty"`
}

func (o UpdateDashboardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDashboardRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDashboardRequestBody", string(data)}, " ")
}
