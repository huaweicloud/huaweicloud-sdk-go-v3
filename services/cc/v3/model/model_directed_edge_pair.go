package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DirectedEdgePair 分支网络两端接入对象。
type DirectedEdgePair struct {

	// 分支网络连接的两个端点定义，可能为两个点，也可能为两个单向边，长度固定为2的数组。
	EdgePair []DirectedEdge `json:"edge_pair"`
}

func (o DirectedEdgePair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectedEdgePair struct{}"
	}

	return strings.Join([]string{"DirectedEdgePair", string(data)}, " ")
}
