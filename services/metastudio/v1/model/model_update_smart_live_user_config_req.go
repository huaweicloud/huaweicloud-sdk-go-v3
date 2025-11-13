package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSmartLiveUserConfigReq 租户直播配置修改请求
type UpdateSmartLiveUserConfigReq struct {
	LiveExitConfig *LiveExitConfig `json:"live_exit_config,omitempty"`

	LiveEventCallbackConfig *LiveEventCallBackConfig `json:"live_event_callback_config,omitempty"`

	LiveNotifyConfig *LiveNotifyConfigReq `json:"live_notify_config,omitempty"`
}

func (o UpdateSmartLiveUserConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSmartLiveUserConfigReq struct{}"
	}

	return strings.Join([]string{"UpdateSmartLiveUserConfigReq", string(data)}, " ")
}
