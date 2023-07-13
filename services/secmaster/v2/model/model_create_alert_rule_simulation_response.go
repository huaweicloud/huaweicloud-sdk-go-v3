package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertRuleSimulationResponse Response Object
type CreateAlertRuleSimulationResponse struct {

	// alert_count
	AlertCount *int32 `json:"alert_count,omitempty"`

	// severity. TIPS, LOW, MEDIUM, HIGH, FATAL
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
