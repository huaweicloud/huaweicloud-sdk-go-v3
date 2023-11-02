package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertRuleSimulationResponse Response Object
type CreateAlertRuleSimulationResponse struct {

	// 告警数量。Alert count.
	AlertCount *int32 `json:"alert_count,omitempty"`

	// 严重程度，提示、低危、中危、高危、致命。Severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity *string `json:"severity,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAlertRuleSimulationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertRuleSimulationResponse struct{}"
	}

	return strings.Join([]string{"CreateAlertRuleSimulationResponse", string(data)}, " ")
}
