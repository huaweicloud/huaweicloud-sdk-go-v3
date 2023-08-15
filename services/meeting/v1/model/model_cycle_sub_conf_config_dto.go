package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CycleSubConfConfigDto struct {

	// 允许加入会议的范围。 - 0: 所有用户 - 2: 企业内用户 - 3: 被邀请用户
	CallInRestriction *int32 `json:"callInRestriction,omitempty"`

	// 允许加入网络研讨会的观众范围。 - 0：所有用户 - 2：企业内用户和被邀请用户
	AudienceCallInRestriction *int32 `json:"audienceCallInRestriction,omitempty"`

	// 是否允许来宾启动会议。 - false:禁止来宾启动会议 - true：允许来宾启动会议 > 仅随机会议ID的会议生效。
	AllowGuestStartConf *bool `json:"allowGuestStartConf,omitempty"`

	// 是否启用等候室。
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
