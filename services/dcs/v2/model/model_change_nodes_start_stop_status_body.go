package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeNodesStartStopStatusBody struct {

	// 实例需要启停的节点ID列表。未配置该参数时，默认启停实例全部节点。
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 对实例节点的操作：  start：开启节点  stop：关闭节点
	Action *string `json:"action,omitempty"`
}

func (o ChangeNodesStartStopStatusBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeNodesStartStopStatusBody struct{}"
	}

	return strings.Join([]string{"ChangeNodesStartStopStatusBody", string(data)}, " ")
}
