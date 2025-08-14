package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFileStatisticRequest Request Object
type ShowFileStatisticRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 开始时间，13位时间戳 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	BeginTime *int64 `json:"begin_time,omitempty"`

	// **参数解释**: 结束时间，13位时间戳 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ShowFileStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileStatisticRequest struct{}"
	}

	return strings.Join([]string{"ShowFileStatisticRequest", string(data)}, " ")
}
