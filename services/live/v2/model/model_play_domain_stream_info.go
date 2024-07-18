package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PlayDomainStreamInfo struct {

	// 播放域名
	PlayDomain *string `json:"play_domain,omitempty"`

	// 播放流名
	Stream *string `json:"stream,omitempty"`

	// 播放的协议，支持flv,hls,rtmp。
	Protocol *string `json:"protocol,omitempty"`

	// 1分钟粒度的带宽值，单位为bps。
	Bandwidth *int64 `json:"bandwidth,omitempty"`

	// 1分钟粒度的在线人数。
	Online *int64 `json:"online,omitempty"`
}

func (o PlayDomainStreamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlayDomainStreamInfo struct{}"
	}

	return strings.Join([]string{"PlayDomainStreamInfo", string(data)}, " ")
}
