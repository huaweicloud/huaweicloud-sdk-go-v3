package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDashboardRequestBody struct {

	// 自定义监控面板名称
	DashboardName *string `json:"dashboard_name,omitempty"`

	// 企业项目Id
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// 监控面板id
	DashboardId *string `json:"dashboard_id,omitempty"`

	// 监控视图展示模式，0表示自定义坐标，1代表每行一个
	RowWidgetNum *int32 `json:"row_widget_num,omitempty"`
}

func (o CreateDashboardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDashboardRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDashboardRequestBody", string(data)}, " ")
}
