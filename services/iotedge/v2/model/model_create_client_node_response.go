package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClientNodeResponse Response Object
type CreateClientNodeResponse struct {

	// 推送通道ID
	ChannelId *string `json:"channel_id,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 路由分配到节点的时间
	AllottedTime *string `json:"allotted_time,omitempty"`

	// 节点实例化后通道的连接和推送信息的修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 下发时间，表示通道是否已经同步到了节点
	SynchronizedTime *string `json:"synchronized_time,omitempty"`

	// 下发状态，表示是否已同步到了节点
	SynchronizedStatus *bool `json:"synchronized_status,omitempty"`
	HttpStatusCode     int   `json:"-"`
}

func (o CreateClientNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClientNodeResponse struct{}"
	}

	return strings.Join([]string{"CreateClientNodeResponse", string(data)}, " ")
}
