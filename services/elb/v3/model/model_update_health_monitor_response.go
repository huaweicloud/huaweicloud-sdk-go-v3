package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHealthMonitorResponse Response Object
type UpdateHealthMonitorResponse struct {

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId *string `json:"request_id,omitempty"`

	Healthmonitor  *HealthMonitor `json:"healthmonitor,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateHealthMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthMonitorResponse struct{}"
	}

	return strings.Join([]string{"UpdateHealthMonitorResponse", string(data)}, " ")
}
