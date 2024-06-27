package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MediaQos 某种媒体的QoS，包括视频、音频或辅流
type MediaQos struct {

	// 客户端-->服务器方向QoS
	UpList *[]Qos `json:"upList,omitempty"`

	// 服务器-->客户端方向QoS
	DownList *[]Qos `json:"downList,omitempty"`
}

func (o MediaQos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MediaQos struct{}"
	}

	return strings.Join([]string{"MediaQos", string(data)}, " ")
}
