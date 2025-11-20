package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmConfigResponse Response Object
type ShowAlarmConfigResponse struct {

	// 0开启 1关闭
	Blackhole *string `json:"blackhole,omitempty"`

	// 0开启 1关闭
	Ddos *string `json:"ddos,omitempty"`

	// topic名称
	TopicName *string `json:"topic_name,omitempty"`

	// topic订阅链接
	TopicUrn       *string `json:"topic_urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmConfigResponse", string(data)}, " ")
}
