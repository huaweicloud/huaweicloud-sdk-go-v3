package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityScoreResponse Response Object
type ShowSecurityScoreResponse struct {

	// **参数解释**: 安全风险总数 **取值范围**: 最小值0，最大值2147483647
	SecurityRiskNum *int32 `json:"security_risk_num,omitempty"`

	// **参数解释**: 安全评分 **取值范围**: 最小值0，最大值100
	SecurityScore *float32 `json:"security_score,omitempty"`

	// **参数解释**: 更新时间，毫秒级时间戳 **取值范围**: 最小值0，最大值9223372036854775807
	DetectTime *int64 `json:"detect_time,omitempty"`

	// **参数解释**: 免费体检主机数 **取值范围**: 最小值0，最大值2147483647
	FreeHealthCheckTotalHostNum *int32 `json:"free_health_check_total_host_num,omitempty"`
	HttpStatusCode              int    `json:"-"`
}

func (o ShowSecurityScoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityScoreResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityScoreResponse", string(data)}, " ")
}
