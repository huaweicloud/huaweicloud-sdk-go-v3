package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Channel struct {

	// 通道名
	Name *string `json:"name,omitempty"`

	// 通道细节描述
	Description *string `json:"description,omitempty"`

	// 通道创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 共识策略
	Consensus *string `json:"consensus,omitempty"`

	// key:组织名，value:节点名称列表
	Peers map[string][]string `json:"peers,omitempty"`

	// key:组织名，value:节点名称列表
	ConsensusNodes map[string][]string `json:"consensusNodes,omitempty"`
}

func (o Channel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Channel struct{}"
	}

	return strings.Join([]string{"Channel", string(data)}, " ")
}
