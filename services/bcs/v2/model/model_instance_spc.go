package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceSpc struct {
	// 单个组织支持的最大peer节点数量

	OrgPeerMaxNum *int64 `json:"org_peer_max_num,omitempty"`
	// 单个联盟链支持的最大order节点数量

	OrdererMaxNum *int64 `json:"orderer_max_num,omitempty"`
	// 单个联盟链支持的最大租户数量

	MemberMaxNum *int64 `json:"member_max_num,omitempty"`
}

func (o InstanceSpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSpc struct{}"
	}

	return strings.Join([]string{"InstanceSpc", string(data)}, " ")
}
