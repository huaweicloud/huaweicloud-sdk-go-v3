package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIefMessageChannelReq 创建IEF消息通道请求body体。
type CreateIefMessageChannelReq struct {

	// 创建IEF消息通道指定的边缘节点ID，uuid格式
	NodeId string `json:"node_id"`
}

func (o CreateIefMessageChannelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIefMessageChannelReq struct{}"
	}

	return strings.Join([]string{"CreateIefMessageChannelReq", string(data)}, " ")
}
