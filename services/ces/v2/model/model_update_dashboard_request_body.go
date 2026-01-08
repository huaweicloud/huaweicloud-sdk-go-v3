package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDashboardRequestBody struct {

	// **参数解释** 自定义监控看板名称 **约束限制** 不涉及 **取值范围** 字符串包含中文字符，字母，数字，下划线（_），横线（-）长度为[1,128]个字符 **默认取值** 不涉及
	DashboardName *string `json:"dashboard_name,omitempty"`

	// **参数解释** 监控看板是否标记收藏 **约束限制** 不涉及 **取值范围** - true 收藏 - false 未收藏 **默认取值** 不涉及
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// **参数解释** 监控视图展示模式 **约束限制** 不涉及 **取值范围** - 0 自定义坐标 - 1 每行一个 - 2 每行两个 - 3 每行三个 - 4 每行四个 **默认取值** 不涉及
	RowWidgetNum int32 `json:"row_widget_num"`

	ExtendInfo *ExtendInfo `json:"extend_info,omitempty"`
}

func (o UpdateDashboardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDashboardRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDashboardRequestBody", string(data)}, " ")
}
