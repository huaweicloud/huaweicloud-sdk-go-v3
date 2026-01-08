package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDashboardWidgetsResponse Response Object
type ListDashboardWidgetsResponse struct {

	// **参数解释** 监控视图列表
	Widgets        *[]WidgetInfoWithId `json:"widgets,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListDashboardWidgetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDashboardWidgetsResponse struct{}"
	}

	return strings.Join([]string{"ListDashboardWidgetsResponse", string(data)}, " ")
}
