package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDashboardWidgetsRequest Request Object
type CreateDashboardWidgetsRequest struct {

	// **参数解释**: 监控看板id **约束限制**: 不涉及。 **取值范围**: 以db开头，包含22个字母和数字，长度为24个字符 **默认取值**: 不涉及。
	DashboardId string `json:"dashboard_id"`

	Body *[]BaseWidgetInfo `json:"body,omitempty"`
}

func (o CreateDashboardWidgetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDashboardWidgetsRequest struct{}"
	}

	return strings.Join([]string{"CreateDashboardWidgetsRequest", string(data)}, " ")
}
