package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestAllowWaitingParticipantReqBody struct {

	// 等候室成员标志。通过监听[[会议级事件推送中的“等候室成员列表信息”](https://support.huaweicloud.com/api-meeting/meeting_21_1507.html)](tag:hws)[[会议级事件推送中的“等候室成员列表信息”](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_1507.html)](tag:hk)事件获得。
	ParticipantID *string `json:"participantID,omitempty"`

	// 允许所有等候者入会。 * false：指定等候者 * true：所有等候者入会
	AllowAll *bool `json:"allowAll,omitempty"`
}

func (o RestAllowWaitingParticipantReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestAllowWaitingParticipantReqBody struct{}"
	}

	return strings.Join([]string{"RestAllowWaitingParticipantReqBody", string(data)}, " ")
}
