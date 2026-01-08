package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDashboardWidgetsRequest Request Object
type ListDashboardWidgetsRequest struct {

	// **参数解释**: 监控看板id **约束限制**: 不涉及。 **取值范围**: 以db开头，包含22个字母和数字，长度为24个字符 **默认取值**: 不涉及。
	DashboardId string `json:"dashboard_id"`

	// **参数解释**: 视图所在的分组id **约束限制**: 不涉及。 **取值范围**: 字符串必须以dg开头，后跟22个字母和数字，总长度为24个字符或者为default，default代表不分组 **默认取值**: 不涉及。
	GroupId *string `json:"group_id,omitempty"`
}

func (o ListDashboardWidgetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDashboardWidgetsRequest struct{}"
	}

	return strings.Join([]string{"ListDashboardWidgetsRequest", string(data)}, " ")
}
