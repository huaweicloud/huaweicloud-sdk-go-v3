package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubscriptionUserResponseDingTalkBotEndpointInfo struct {

	// 钉钉企业用户的userId。
	Endpoint *string `json:"endpoint,omitempty"`

	// 钉钉创建的机器人编码。
	RobotCode *string `json:"robot_code,omitempty"`
}

func (o ListSubscriptionUserResponseDingTalkBotEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionUserResponseDingTalkBotEndpointInfo struct{}"
	}

	return strings.Join([]string{"ListSubscriptionUserResponseDingTalkBotEndpointInfo", string(data)}, " ")
}
