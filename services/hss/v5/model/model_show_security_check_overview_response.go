package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityCheckOverviewResponse Response Object
type ShowSecurityCheckOverviewResponse struct {

	// **参数解释**: 风险评分 **取值范围**: 最小值0，最大值100
	RiskRating *int32 `json:"risk_rating,omitempty"`

	// **参数解释**: 风险级别 **取值范围**: - Security：安全 - Low：低危 - Medium：中危 - High：高危
	Severity *string `json:"severity,omitempty"`

	// **参数解释**: 最新检测时间，毫秒级时间戳 **取值范围**: 最小值0，最大值9223372036854775807（UTC时区，从1970-01-01 00:00:00开始计算）
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**: 防护的服务器数量 **取值范围**: 最小值0，最大值2000000
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**: 安全风险数量 **取值范围**: 最小值0，最大值2147483647
	RiskNum *int32 `json:"risk_num,omitempty"`

	// **参数解释**: 定时体检开启状态 **取值范围**: - true：已开启 - false：未开启
	ScheduledCheckStatus *bool `json:"scheduled_check_status,omitempty"`
	HttpStatusCode       int   `json:"-"`
}

func (o ShowSecurityCheckOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityCheckOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityCheckOverviewResponse", string(data)}, " ")
}
