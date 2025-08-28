package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBotMScoreDistributionRequest Request Object
type ListBotMScoreDistributionRequest struct {

	// **参数解释：** 开始时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	StartTime int64 `json:"start_time"`

	// **参数解释：** 结束时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EndTime int64 `json:"end_time"`

	// **参数解释：** 租户Id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释：** 企业项目Id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 域名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Hosts *[]string `json:"hosts,omitempty"`

	// **参数解释：** 域名,查询的域名列表，与hosts二选一，如果hosts不为空，以hosts为准 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Domains *[]string `json:"domains,omitempty"`

	// **参数解释：** 区域 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释：** 站点 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Site *string `json:"site,omitempty"`
}

func (o ListBotMScoreDistributionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBotMScoreDistributionRequest struct{}"
	}

	return strings.Join([]string{"ListBotMScoreDistributionRequest", string(data)}, " ")
}
