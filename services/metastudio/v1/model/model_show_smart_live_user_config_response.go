package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSmartLiveUserConfigResponse Response Object
type ShowSmartLiveUserConfigResponse struct {
	LiveExitConfig *LiveExitConfig `json:"live_exit_config,omitempty"`

	LiveEventCallbackConfig *LiveEventCallBackConfig `json:"live_event_callback_config,omitempty"`

	LiveNotifyConfig *LiveNotifyConfig `json:"live_notify_config,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSmartLiveUserConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSmartLiveUserConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowSmartLiveUserConfigResponse", string(data)}, " ")
}
