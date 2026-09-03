package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubscriptionUserRequestDingTalkBotEndpointInfo struct {

	// 钉钉企业用户的userId。
	Endpoint string `json:"endpoint"`

	// 个人钉钉appKey字段。
	AppKey string `json:"app_key"`

	// 个人钉钉appSecret字段。
	AppSecret string `json:"app_secret"`

	// 个人钉钉robotCode字段。
	RobotCode string `json:"robot_code"`
}

func (o CreateSubscriptionUserRequestDingTalkBotEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionUserRequestDingTalkBotEndpointInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionUserRequestDingTalkBotEndpointInfo", string(data)}, " ")
}
