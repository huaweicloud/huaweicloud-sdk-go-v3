package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesClient 客户端。
type PoliciesClient struct {

	// 自动重连间隔（秒）。取值范围为[1-50]。默认：5。
	AutomaticReconnectionInterval *int32 `json:"automatic_reconnection_interval,omitempty"`

	// 自动重连会话保持时长（秒）。取值范围为[0-180]。默认：180。
	SessionPersistenceTime *int32 `json:"session_persistence_time,omitempty"`

	// 锁屏后自动关闭本地显示器。取值为： false：表示关闭。 true：表示开启。
	AutocloseMonitorAfterLocked *bool `json:"autoclose_monitor_after_locked,omitempty"`

	AutocloseMonitorOptions *PoliciesClientAutocloseMonitorOptions `json:"autoclose_monitor_options,omitempty"`

	// 防截屏策略开关。 false：表示关闭。 true：表示开启。
	ForbidScreenCapture *bool `json:"forbid_screen_capture,omitempty"`

	// 终端机器加域校验开关。 false：表示关闭。 true：表示开启。
	ClientMachineJoinDomain *bool `json:"client_machine_join_domain,omitempty"`

	ClientType *PoliciesClientClientType `json:"client_type,omitempty"`
}

func (o PoliciesClient) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesClient struct{}"
	}

	return strings.Join([]string{"PoliciesClient", string(data)}, " ")
}
