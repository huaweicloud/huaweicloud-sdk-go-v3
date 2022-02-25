package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShrinkNodesReq struct {
	// 需要下线的节点ID。

	ShrinkNodes []string `json:"shrinkNodes"`
}

func (o ShrinkNodesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkNodesReq struct{}"
	}

	return strings.Join([]string{"ShrinkNodesReq", string(data)}, " ")
}
