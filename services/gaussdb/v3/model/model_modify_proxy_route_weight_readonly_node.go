package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyProxyRouteWeightReadonlyNode 只读节点权重配置信息。
type ModifyProxyRouteWeightReadonlyNode struct {

	// 只读节点id
	Id string `json:"id"`

	// 只读节点权重： - 如果路由模式为0，取值为0~1000； - 如果路由模式为1或2，取值为0或1。
	Weight int32 `json:"weight"`
}

func (o ModifyProxyRouteWeightReadonlyNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyProxyRouteWeightReadonlyNode struct{}"
	}

	return strings.Join([]string{"ModifyProxyRouteWeightReadonlyNode", string(data)}, " ")
}
