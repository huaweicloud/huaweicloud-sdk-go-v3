package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChannelCreateReqPolicy struct {

	// 固定值：allow_account_to_put_events
	Sid *string `json:"Sid,omitempty"`

	// 固定值：Allow
	Effect *string `json:"Effect,omitempty"`

	// domain白名单
	Principal map[string]ChannelCreateReqPrincipal `json:"Principal,omitempty"`

	// 固定值：eg:channels:putEvents
	Action *string `json:"Action,omitempty"`

	// 固定格式：urn:eg:cn-east-2:{project_id}:channel:{channel_name}
	Resource *string `json:"Resource,omitempty"`
}

func (o ChannelCreateReqPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelCreateReqPolicy struct{}"
	}

	return strings.Join([]string{"ChannelCreateReqPolicy", string(data)}, " ")
}
