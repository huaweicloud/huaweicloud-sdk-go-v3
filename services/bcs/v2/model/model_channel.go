package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Channel struct {

	// 通道名
	Name *string `json:"name,omitempty" xml:"name"`

	// 通道细节描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 通道创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 共识策略
	Consensus *string `json:"consensus,omitempty" xml:"consensus"`

	// key:组织名，value:节点名称列表
	Peers map[string][]string `json:"peers,omitempty" xml:"peers"`

	// key:组织名，value:节点名称列表
	ConsensusNodes map[string][]string `json:"consensusNodes,omitempty" xml:"consensusNodes"`
}

func (o Channel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Channel struct{}"
	}

	return strings.Join([]string{"Channel", string(data)}, " ")
}
