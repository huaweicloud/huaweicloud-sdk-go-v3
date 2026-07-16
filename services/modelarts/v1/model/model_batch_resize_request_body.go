package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchResizeRequestBody struct {

	// **参数解释**：扩缩容的超节点批次信息。 **约束限制**：单次最多50个超节点。
	Nodes []ResizeNodeInfo `json:"nodes"`

	Source *NodeResizeParams `json:"source"`

	Target *NodeResizeParams `json:"target"`
}

func (o BatchResizeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResizeRequestBody struct{}"
	}

	return strings.Join([]string{"BatchResizeRequestBody", string(data)}, " ")
}
