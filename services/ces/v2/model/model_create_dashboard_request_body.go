package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDashboardRequestBody struct {

	// **参数描述**： 自定义监控看板名称 **约束限制** 不涉及 **取值范围** 长度为[1,128]个字符，只允许中文、英文、数字0-9、_和- **默认取值** 不涉及
	DashboardName *string `json:"dashboard_name,omitempty"`

	// **参数解释** 企业项目ID **约束限制** 不涉及 **取值范围** 只能包含小写字母、数字、“-”、“_”，可以自定义企业项目ID，长度为36个字符。也可以为0（代表默认企业项目ID）。 **默认取值** 不涉及
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// **参数描述**： 监控看板id **约束限制** 不涉及 **取值范围** 以db开头，包含22个字母和数字，长度为24个字符 **默认取值** 不涉及
	DashboardId *string `json:"dashboard_id,omitempty"`

	// **参数描述**： 监控视图展示模式 **约束限制** 不涉及 **取值范围** - 0:表示自定义坐标 - 1:代表每行1个监控视图 - 2:代表每行2个监控视图 - 3:代表每行3个监控视图 **默认取值** 3
	RowWidgetNum *int32 `json:"row_widget_num,omitempty"`
}

func (o CreateDashboardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDashboardRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDashboardRequestBody", string(data)}, " ")
}
