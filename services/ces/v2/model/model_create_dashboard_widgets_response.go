package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDashboardWidgetsResponse Response Object
type CreateDashboardWidgetsResponse struct {

	// **参数解释** 批量创建监控视图返回结果
	WidgetIds      *[]string `json:"widget_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateDashboardWidgetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDashboardWidgetsResponse struct{}"
	}

	return strings.Join([]string{"CreateDashboardWidgetsResponse", string(data)}, " ")
}
