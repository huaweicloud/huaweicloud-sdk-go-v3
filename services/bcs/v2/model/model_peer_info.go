package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeerInfo struct {
	// 组织名称

	Name *string `json:"name,omitempty"`
	// 节点数量

	NodeCnt *int64 `json:"node_cnt,omitempty"`
	// 节点状态，分为创建中（IsCreating），升级中（IsUpgrading），扩缩容中（Adding/IsScaling），删除中（Isdeleting），正常（Normal），异常（AbNormal），未知（其余值）

	Status *string `json:"status,omitempty"`
	// 节点状态，形式如：1/1，分母是该组织下节点总数，分子是正常节点个数

	StatusDetail *string `json:"status_detail,omitempty"`
	// 节点对应pvc名称

	PvcName *string `json:"pvc_name,omitempty"`
	// Peer节点域名/IP地址

	Address *[]PeerAddress `json:"address,omitempty"`
}

func (o PeerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeerInfo struct{}"
	}

	return strings.Join([]string{"PeerInfo", string(data)}, " ")
}
