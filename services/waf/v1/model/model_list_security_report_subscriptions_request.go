package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityReportSubscriptionsRequest Request Object
type ListSecurityReportSubscriptionsRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。仅支持查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - all_granted_eps：所有企业项目  **默认取值：** all_granted_eps
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// **参数解释：** 报告模板名称 **约束限制：** 不涉及 **取值范围：** 只能由中文、字母、数字和括号内所列符号（_-.:：）组成，且不能超过256个字符长度。 **默认取值：** 不涉及
	ReportName *string `json:"report_name,omitempty"`

	// **参数解释：** 报告类型 **约束限制：** 不涉及 **取值范围：** - daily_report：安全日报 - weekly_report：安全周报 - monthly_report：安全月报 - custom_report：自定义报告  **默认取值：** 不涉及
	ReportCategory *string `json:"report_category,omitempty"`

	// **参数解释：** 开启状态 **约束限制：** 不涉及 **取值范围：** - opened：已开启 - closed：已关闭  **默认取值：** 不涉及
	ReportStatus *string `json:"report_status,omitempty"`

	// **参数解释：** 分页查询的起始位置，表示从第几条记录开始返回（从0开始计数）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 分页查询的单页返回数量，控制每次请求返回的记录条数。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 1000
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSecurityReportSubscriptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityReportSubscriptionsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityReportSubscriptionsRequest", string(data)}, " ")
}
