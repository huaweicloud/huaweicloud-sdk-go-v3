package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIefMessageChannelRequestBody 创建IEF消息通道请求body体。
type CreateIefMessageChannelRequestBody struct {

	// 创建IEF消息通道指定的边缘节点ID，uuid格式
	NodeId string `json:"node_id"`
}

func (o CreateIefMessageChannelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIefMessageChannelRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIefMessageChannelRequestBody", string(data)}, " ")
}
