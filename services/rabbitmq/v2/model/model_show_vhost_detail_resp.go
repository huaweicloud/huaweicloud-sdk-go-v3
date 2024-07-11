package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVhostDetailResp struct {

	// Vhost名称。
	Name *string `json:"name,omitempty"`

	// 是否开启消息轨迹[（AMQP版本不涉及此字段）](tag:hws,hws_hk)。
	Tracing *bool `json:"tracing,omitempty"`
}

func (o ShowVhostDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVhostDetailResp struct{}"
	}

	return strings.Join([]string{"ShowVhostDetailResp", string(data)}, " ")
}
