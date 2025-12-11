package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceMonitorExtendResponse Response Object
type ShowInstanceMonitorExtendResponse struct {

	// 实例秒级监控开关。  - true，表示开启。 - false，表示关闭。
	MonitorSwitch *bool `json:"monitor_switch,omitempty"`

	// 采集周期，仅在monitor_switch为true时返回。  - 1：采集周期为1s。 - 5：采集周期为5s。
	Period *int32 `json:"period,omitempty"`

	// **参数解释**：  能否开启秒级监控。  **取值范围**： - true: 可以开启。 - false：不能开启，8vcpu以下规格不允许开启秒级监控。
	Allow          *bool `json:"allow,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowInstanceMonitorExtendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMonitorExtendResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceMonitorExtendResponse", string(data)}, " ")
}
