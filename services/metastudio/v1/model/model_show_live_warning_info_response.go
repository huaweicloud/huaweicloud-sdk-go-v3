package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLiveWarningInfoResponse Response Object
type ShowLiveWarningInfoResponse struct {

	// 开播风险告警列表。
	LiveWarningInfo *[]LiveWarningItem `json:"live_warning_info,omitempty"`

	// **参数解释**： 配置的最大直播时长。单位小时。 0 为不限制。 **约束限制**： 停止直播逻辑配置为立即停止则直播停止误差在5分钟之内。其他逻辑则加上处理时长。 **默认取值**： 不设置则表示不限时。
	LimitDuration *int32 `json:"limit_duration,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLiveWarningInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLiveWarningInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowLiveWarningInfoResponse", string(data)}, " ")
}
