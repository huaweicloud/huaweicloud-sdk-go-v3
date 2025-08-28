package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHealthMonitorRequest Request Object
type DeleteHealthMonitorRequest struct {

	// **参数解释**：健康检查ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o DeleteHealthMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"DeleteHealthMonitorRequest", string(data)}, " ")
}
