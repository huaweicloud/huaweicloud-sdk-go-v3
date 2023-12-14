package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestSetAttendeeLanChannelBody 传译员信息
type RestSetAttendeeLanChannelBody struct {

	// 会场标识列表。
	ParticipantIDs *[]string `json:"participantIDs,omitempty"`

	// 与会者收听的语言频道，普通与会者听与说一致。
	ListenLanChannel string `json:"listenLanChannel"`

	// 与会者发言的语言频道，普通与会者听与说一致。
	SpeakLanChannel string `json:"speakLanChannel"`

	// 是否包含原声，0：不包含，1：包含。
	IncludeOriginalVoice *int32 `json:"includeOriginalVoice,omitempty"`
}

func (o RestSetAttendeeLanChannelBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSetAttendeeLanChannelBody struct{}"
	}

	return strings.Join([]string{"RestSetAttendeeLanChannelBody", string(data)}, " ")
}
