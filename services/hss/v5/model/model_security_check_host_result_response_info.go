package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckHostResultResponseInfo 服务器的安全体检结果信息
type SecurityCheckHostResultResponseInfo struct {
	HostInfo *SecurityCheckHostInfo `json:"host_info,omitempty"`

	// **参数解释**： 风险级别 **取值范围**： - high：高危 - medium：中危 - low：低危 - safe：安全，无风险
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 风险评分 **取值范围**： 不涉及
	RiskRating *int32 `json:"risk_rating,omitempty"`

	// **参数解释**： 风险数量 **取值范围**： 不涉及
	RiskNum *int32 `json:"risk_num,omitempty"`

	// **参数解释**： 最新检测时间 **取值范围**： 不涉及
	ScanTime *int64 `json:"scan_time,omitempty"`
}

func (o SecurityCheckHostResultResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckHostResultResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckHostResultResponseInfo", string(data)}, " ")
}
