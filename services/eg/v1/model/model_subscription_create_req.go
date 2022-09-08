package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionCreateReq struct {

	// 事件订阅名称，租户下唯一，由字母、数字、点、下划线和中划线组成，必须字母或数字开头，长度为1~128字符。
	Name string `json:"name"`

	// 事件订阅描述
	Description *string `json:"description,omitempty"`

	// 所属事件通道ID
	ChannelId string `json:"channel_id"`

	// 事件订阅的事件源列表，当前仅支持订阅一个事件源。
	Sources []SubscriptionSource `json:"sources"`

	// 事件目标列表，至少订阅一个，最多五个事件目标。
	Targets []SubscriptionTarget `json:"targets"`
}

func (o SubscriptionCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionCreateReq struct{}"
	}

	return strings.Join([]string{"SubscriptionCreateReq", string(data)}, " ")
}
