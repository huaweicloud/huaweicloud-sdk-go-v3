package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogsRequest Request Object
type ShowLogsRequest struct {

	// **参数解释：** 域名 **约束限制：** 只支持单个域名，如：www.test1.com **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainName string `json:"domain_name"`

	// **参数解释：** 查询开始时间 **约束限制：** 不涉及 **取值范围：** 时间格式为整点毫秒时间戳 **默认取值：** 当天0点
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释：** 查询结束时间 **约束限制：** 不涉及 **取值范围：** - 不包含结束时间 - 与开始时间的最大跨度为30天 - 时间格式为整点毫秒时间戳 **默认取值：** 开始时间加1天
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释：** 查询时单页数量 **约束限制：** 不涉及 **取值范围：** 1-10000 **默认取值：** 10
	PageSize *int32 `json:"page_size,omitempty"`

	// **参数解释：** 当前查询第几页 **约束限制：** 不涉及 **取值范围：** 1-65535 **默认取值：** 1
	PageNumber *int32 `json:"page_number,omitempty"`

	// **参数解释：** 企业项目id > 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id  **约束限制：** - 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目 - 当使用子账号调用接口时，该参数必传 **取值范围：** all表示所有项目 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogsRequest struct{}"
	}

	return strings.Join([]string{"ShowLogsRequest", string(data)}, " ")
}
