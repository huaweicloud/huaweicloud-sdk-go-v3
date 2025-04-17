package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DashBoardInfo struct {

	// 监控看板id
	DashboardId *string `json:"dashboard_id,omitempty"`

	// 自定义监控看板名称
	DashboardName *string `json:"dashboard_name,omitempty"`

	// 企业项目Id
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// 监控视图展示模式，0表示自定义坐标，1代表每行一个
	RowWidgetNum *int32 `json:"row_widget_num,omitempty"`

	// 监控看板是否标记收藏, true: 收藏, false: 未收藏
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// 监控看板的创建用户名
	CreatorName *string `json:"creator_name,omitempty"`

	// 监控看板创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 看板下的视图总数
	WidgetsNum *int32 `json:"widgets_num,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 子产品标识
	SubProduct *string `json:"sub_product,omitempty"`

	// 监控大盘模板id
	DashboardTemplateId *string `json:"dashboard_template_id,omitempty"`
}

func (o DashBoardInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DashBoardInfo struct{}"
	}

	return strings.Join([]string{"DashBoardInfo", string(data)}, " ")
}
