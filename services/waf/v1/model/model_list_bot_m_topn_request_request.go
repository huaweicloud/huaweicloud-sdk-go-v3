package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBotMTopnRequestRequest Request Object
type ListBotMTopnRequestRequest struct {

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

	// **参数解释：** n的取值 **约束限制：** 不涉及 **取值范围：** 1-20 **默认取值：** 5
	Topn *int64 `json:"topn,omitempty"`

	// **参数解释：** 区域 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释：** 站点 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Site *string `json:"site,omitempty"`
}

func (o ListBotMTopnRequestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBotMTopnRequestRequest struct{}"
	}

	return strings.Join([]string{"ListBotMTopnRequestRequest", string(data)}, " ")
}
