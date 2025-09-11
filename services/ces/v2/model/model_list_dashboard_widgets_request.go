package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDashboardWidgetsRequest Request Object
type ListDashboardWidgetsRequest struct {

	// **参数解释**: 监控看板id **约束限制**: 不涉及。 **取值范围**: 以db开头，包含22个字母和数字，长度为24个字符 **默认取值**: 不涉及。
	DashboardId string `json:"dashboard_id"`

	// 视图所在的分组id
	GroupId *string `json:"group_id,omitempty"`
}

func (o ListDashboardWidgetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDashboardWidgetsRequest struct{}"
	}

	return strings.Join([]string{"ListDashboardWidgetsRequest", string(data)}, " ")
}
