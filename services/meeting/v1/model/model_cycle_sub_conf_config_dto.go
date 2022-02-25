package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CycleSubConfConfigDto struct {
	// |参数名称：允许呼入的范围| |参数描述：允许呼入的范围。 0: 所有用户 1: 非匿名用户（手机pstn入会视为匿名入会） 2: 企业内用户 3: 被邀请用户| |取值范围：[0,3]|

	CallInRestriction *int32 `json:"callInRestriction,omitempty"`
	// |参数名称：网络研讨会观众允许呼入的范围| |参数描述：允许呼入的范围。 0: 所有用户 2: 企业内用户和被邀请用户|

	AudienceCallInRestriction *int32 `json:"audienceCallInRestriction,omitempty"`
	// 参数名称：是否允许来宾启动会议(随机会议) false:禁止来宾启动会议 true：允许来宾启动会议

	AllowGuestStartConf *bool `json:"allowGuestStartConf,omitempty"`
	// 是否启用等候室

	EnableWaitingRoom *bool `json:"enableWaitingRoom,omitempty"`

	ShowAudienceCountInfo *ShowAudienceCountInfo `json:"showAudienceCountInfo,omitempty"`
}

func (o CycleSubConfConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CycleSubConfConfigDto struct{}"
	}

	return strings.Join([]string{"CycleSubConfConfigDto", string(data)}, " ")
}
